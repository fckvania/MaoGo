package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"mao/src/typings"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

var youtubeRegex = regexp.MustCompile(`^(https?://)?(www\.)?(youtube\.com|youtu\.be)/.+`)

func IsYoutubeURL(url string) bool {
	return youtubeRegex.MatchString(url)
}

func parseSeconds(s string) string {
	seconds, err := strconv.ParseFloat(s, 64)
	if nil != err {
		return ""
	}
	return fmt.Sprintf("%v", seconds)
}

func YoutubeDL(uri string) (typings.YoutubeInfos, error) {
	if !IsYoutubeURL(uri) {
		return typings.YoutubeInfos{}, errors.New("Url Invalid")
	}
	params := url.Values{}
	params.Add("q", uri)
	params.Add("vt", "home")

	req, err := http.PostForm("https://yt1s.com/api/ajaxSearch/index", params)
	if err != nil {
		return typings.YoutubeInfos{}, err
	}

	req.Header.Add("User-Agent", "WhatsApp/2.2353.59")

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return typings.YoutubeInfos{}, err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return typings.YoutubeInfos{}, err
	}

	// Extract information from data map

	info := typings.YoutubeInfo{
		Title:    data["title"].(string),
		Duration: data["t"].(float64),
		Author:   data["a"].(string),
	}
	link := typings.YoutubeLinks{}

	for _, dat := range data["links"].(map[string]interface{})["mp3"].(map[string]interface{}) {
		a := dat.(map[string]interface{})
		if a["f"].(string) != "mp3" {
			continue
		}
		link.Audio = append(link.Audio, typings.YoutubeAV{
			Size:    a["size"].(string),
			Format:  a["f"].(string),
			Quality: a["q"].(string),
			Url: func() (string, error) {
				return Download(data["vid"].(string), a["k"].(string))
			},
		})
	}

	for _, dat := range data["links"].(map[string]interface{})["mp4"].(map[string]interface{}) {
		a := dat.(map[string]interface{})
		if a["f"].(string) != "mp4" {
			continue
		}
		link.Video = append(link.Video, typings.YoutubeAV{
			Size:    a["size"].(string),
			Format:  a["f"].(string),
			Quality: a["q"].(string),
			Url: func() (string, error) {
				return Download(data["vid"].(string), a["k"].(string))
			},
		})
	}

	return typings.YoutubeInfos{
		Info: info,
		Link: link,
	}, nil
}

func Download(id string, k string) (string, error) {
	params := url.Values{}
	params.Add("vid", id)
	params.Add("k", k)

	req, err := http.PostForm("https://yt1s.com/api/ajaxConvert/convert", params)
	if err != nil {
		return "", err
	}

	req.Header.Set("User-Agent", "WhatsApp/2.2353.59")

	defer req.Body.Close()

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return "", err
	}

	var data map[string]interface{}
	err = json.Unmarshal(body, &data)
	if err != nil {
		return "", err
	}

	if data["dlink"] == nil {
		return "", errors.New("Terjadi Kesalahan.")
	}

	return data["dlink"].(string), nil
}
