package api

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"strings"
)

var tiktokRegexp = regexp.MustCompile(`(https:\/\/vm\.tiktok\.com\/[A-Za-z0-9]+)|(https:\/\/vt\.tiktok\.com\/[A-Za-z0-9]+)|(https:\/\/www\.tiktok\.com\/@[^\/]+\/video\/[0-9]+)`)

func isTiktokUrl(url string) bool {
	return tiktokRegexp.MatchString(url)
}

func GetTiktokVideo(url string) (string, error) {
	if !isTiktokUrl(url) {
		return "", errors.New("Url Invalid")
	}
	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}

	resp, err := client.Do(req)
	if err != nil {
		return "", nil
	}
	defer resp.Body.Close()
	location := resp.Header.Get("location")
	data := strings.Split(location, "/")
	data = data[5:]

	if len(data) > 2 {
		return fmt.Sprintf("https://www.tikwm.com/video/media/play/%s.mp4", data[1]), nil
	}

	return fmt.Sprintf("https://www.tikwm.com/video/media/play/%s.mp4", data[0][:19]), nil
}
