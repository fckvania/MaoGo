package api

import (
	"bytes"
	"encoding/json"
	"errors"
	"mime/multipart"
	"net/http"

	"github.com/google/uuid"
)

func MicrosoftDesigner(prompt string) (map[string]interface{}, error) {
	// Generate session ID
	sessionId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	headers := map[string]string{
		"SessionId":         sessionId.String(),
		"X-UserSessionId":   sessionId.String(),
		"Platform":          "Android",
		"Caller":            "DesignerApp",
		"AudienceGroup":     "Production",
		"ClientName":        "DesignerApp",
		"HostApp":           "DesignerApp",
		"User-Agent":        "com.microsoft.designer/2023335701 (Linux; U; Android 13; in_ID; Redmi 9; Build/TQ2A.230505.002.A1; Cronet/114.0.5735.33)",
		"Accept-Encoding":   "gzip, deflate",
		"Transfer-Encoding": "chunked",
		"Authorization":     "", //Get From https://designer.microsoft.com/
	}

	if headers["Authorization"] == "" {
		return nil, errors.New("Need Authorization")
	}

	// Create form data
	formData := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(formData)

	writer.WriteField("dalle-caption", prompt)
	writer.WriteField("dalle-batch-size", "4")
	writer.WriteField("dalle-seed", "744")
	writer.WriteField("dalle-scenario-name", "TextToImage")
	writer.WriteField("dalle-image-response-format", "UrlWithBase64Thumbnail")
	writer.Close()

	// Create HTTP request
	req, err := http.NewRequest("POST", "https://designerapp.officeapps.live.com/designerapp/DallE.ashx?action=GetDallEImagesCogSci", formData)
	if err != nil {
		return nil, err
	}

	// Set headers
	req.Header.Set("Content-Type", writer.FormDataContentType())
	for key, value := range headers {
		req.Header.Set(key, value)
	}

	// Send request and handle response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	// Parse response data
	var data map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, err
}
