package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"mao/src/typings"
	"net/http"
	"strings"
)

var (
	BASE_URL string = "https://sticker-api.openwa.dev/"
)

type sApi struct {
	Build func() []byte
}

func StickerApi(s *typings.Sticker, d *typings.MetadataSticker) *sApi {
	return &sApi{
		Build: func() []byte {
			var path string
			var data string
			data += `"sessionInfo":{"WA_VERSION":"2.2106.5","PAGE_UA":"WhatsApp/2.2037.6 Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36","WA_AUTOMATE_VERSION":"3.6.10 UPDATE AVAILABLE: 3.6.11","BROWSER_VERSION":"HeadlessChrome/88.0.4324.190","OS":"Windows Server 2016","START_TS":1614310326309,"NUM":"6247","LAUNCH_TIME_MS":7934,"PHONE_VERSION":"2.20.205.16"},"config":{"sessionId":"session","headless":true,"qrTimeout":20,"authTimeout":0,"cacheEnabled":false,"useChrome":true,"killProcessOnBrowserClose":true,"throwErrorOnTosBlock":false,"chromiumArgs":["--no-sandbox","--disable-setuid-sandbox","--aggressive-cache-discard","--disable-cache","--disable-application-cache","--disable-offline-load-stale-cache","--disk-cache-size=0"],"executablePath":"C:\\\\Program Files (x86)\\\\Google\\\\Chrome\\\\Application\\\\chrome.exe","skipBrokenMethodsCheck":true,"stickerServerEndpoint":true},`
			data += fmt.Sprintf(`"stickerMetadata": { "author": "%s", "pack": "%s", "keepScale": %v, "removebg": "%v", "circle": %v },`, d.Author, d.Pack, d.KeepScale, d.Removebg, d.Circle)
			if s.Tipe == typings.IMAGE {
				path = "prepareWebp"
				data += `"image": "data:` + http.DetectContentType(s.File) + `;base64,` + base64.StdEncoding.EncodeToString(s.File) + `"`
			} else if s.Tipe == typings.VIDEO {
				path = "convertMp4BufferToWebpDataUrl"
				data += `"processOptions": {"crop":false,"fps":10,"startTime":"00:00:00.0","endTime":"00:00:7.0","loop":0},`
				data += `"file": "data:` + http.DetectContentType(s.File) + `;base64,` + base64.StdEncoding.EncodeToString(s.File) + `"`
			}

			client := &http.Client{}
			req, err := http.NewRequest("POST", BASE_URL+path, bytes.NewBuffer([]byte(`{ `+data+`}`)))
			if err != nil {
				fmt.Println(err)
				return nil
			}
			req.Header.Add("Accept", "application/json, text/plain, /")
			req.Header.Add("Content-Type", "application/json;charset=utf-8")
			req.Header.Add("User-Agent", "WhatsApp/2.2037.6 Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/85.0.4183.83 Safari/537.36")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			defer resp.Body.Close()

			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				fmt.Println(err)
				return nil
			}

			var datap string
			if s.Tipe == typings.IMAGE {
				var result map[string]any
				if err := json.Unmarshal(body, &result); err != nil {
					fmt.Println(err)
					return nil
				}
				datap = fmt.Sprintf("%v", result["webpBase64"])
			} else {
				datap = strings.Split(string(body), ";base64,")[1]
			}
			byteBuffer, _ := base64.StdEncoding.DecodeString(datap)
			return byteBuffer
		},
	}
}
