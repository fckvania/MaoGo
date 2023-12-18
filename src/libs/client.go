package libs

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"strings"

	"go.mau.fi/whatsmeow"
	waProto "go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"google.golang.org/protobuf/proto"
)

func NewClient(client *whatsmeow.Client) *NewClientImpl {
	return &NewClientImpl{
		WA: client,
	}
}

func (client *NewClientImpl) SendText(from types.JID, txt string, opts *waProto.ContextInfo) (whatsmeow.SendResponse, error) {
	ok, er := client.WA.SendMessage(context.Background(), from, &waProto.Message{
		ExtendedTextMessage: &waProto.ExtendedTextMessage{
			Text:        proto.String(txt),
			ContextInfo: opts,
		},
	})
	if er != nil {
		return whatsmeow.SendResponse{}, er
	}
	return ok, nil
}

func (client *NewClientImpl) SendWithNewsLestter(from types.JID, text string, newjid string, newserver int32, name string, opts *waProto.ContextInfo) (whatsmeow.SendResponse, error) {
	ok, er := client.SendText(from, text, &waProto.ContextInfo{
		ForwardedNewsletterMessageInfo: &waProto.ForwardedNewsletterMessageInfo{
			NewsletterJid:     proto.String(newjid),
			NewsletterName:    proto.String(name),
			ServerMessageId:   proto.Int32(newserver),
			ContentType:       waProto.ForwardedNewsletterMessageInfo_UPDATE.Enum(),
			AccessibilityText: proto.String(""),
		},
		IsForwarded:   proto.Bool(true),
		StanzaId:      opts.StanzaId,
		Participant:   opts.Participant,
		QuotedMessage: opts.QuotedMessage,
	})

	if er != nil {
		return whatsmeow.SendResponse{}, er
	}
	return ok, nil
}

func (client *NewClientImpl) SendImage(from types.JID, data []byte, caption string, opts *waProto.ContextInfo) (whatsmeow.SendResponse, error) {
	uploaded, err := client.WA.Upload(context.Background(), data, whatsmeow.MediaImage)
	if err != nil {
		fmt.Printf("Failed to upload file: %v\n", err)
		return whatsmeow.SendResponse{}, err
	}
	resultImg := &waProto.Message{
		ImageMessage: &waProto.ImageMessage{
			Url:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Caption:       proto.String(caption),
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSha256: uploaded.FileEncSHA256,
			FileSha256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			ContextInfo:   opts,
		},
	}
	ok, _ := client.WA.SendMessage(context.Background(), from, resultImg)
	return ok, nil
}

func (client *NewClientImpl) SendVideo(from types.JID, data []byte, caption string, opts *waProto.ContextInfo) (whatsmeow.SendResponse, error) {
	uploaded, err := client.WA.Upload(context.Background(), data, whatsmeow.MediaVideo)
	if err != nil {
		fmt.Printf("Failed to upload file: %v\n", err)
		return whatsmeow.SendResponse{}, err
	}
	resultVideo := &waProto.Message{
		VideoMessage: &waProto.VideoMessage{
			Url:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Caption:       proto.String(caption),
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSha256: uploaded.FileEncSHA256,
			FileSha256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			ContextInfo:   opts,
		},
	}
	ok, er := client.WA.SendMessage(context.Background(), from, resultVideo)
	if er != nil {
		return whatsmeow.SendResponse{}, er
	}
	return ok, nil
}

func (client *NewClientImpl) SendDocument(from types.JID, data []byte, fileName string, caption string, opts *waProto.ContextInfo) (whatsmeow.SendResponse, error) {
	uploaded, err := client.WA.Upload(context.Background(), data, whatsmeow.MediaDocument)
	if err != nil {
		fmt.Printf("Failed to upload file: %v\n", err)
		return whatsmeow.SendResponse{}, err
	}
	resultDoc := &waProto.Message{
		DocumentMessage: &waProto.DocumentMessage{
			Url:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			FileName:      proto.String(fileName),
			Caption:       proto.String(caption),
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSha256: uploaded.FileEncSHA256,
			FileSha256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			ContextInfo:   opts,
		},
	}
	ok, er := client.WA.SendMessage(context.Background(), from, resultDoc)
	if er != nil {
		return whatsmeow.SendResponse{}, er
	}
	return ok, nil
}

func (client *NewClientImpl) UploadImage(data []byte) (string, error) {
	bodyy := &bytes.Buffer{}
	writer := multipart.NewWriter(bodyy)
	part, _ := writer.CreateFormFile("file", "file")
	_, err := io.Copy(part, bytes.NewBuffer(data))
	if err != nil {
		return "", err
	}
	writer.Close()

	// Create request
	req, err := http.NewRequest("POST", "https://telegra.ph/upload", bodyy)
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", writer.FormDataContentType())

	// Send request and handle response
	htt := &http.Client{}
	resp, err := htt.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("HTTP Error: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var uploads []struct {
		Path string `json:"src"`
	}
	if err := json.Unmarshal(body, &uploads); err != nil {
		m := map[string]string{}
		if err := json.Unmarshal(data, &m); err != nil {
			return "", err
		}
		return "", fmt.Errorf("telegraph: %s", m["error"])
	}

	return "https://telegra.ph/" + uploads[0].Path, nil
}

func (client *NewClientImpl) ParseJID(arg string) (types.JID, bool) {
	if arg[0] == '+' {
		arg = arg[1:]
	}
	if !strings.ContainsRune(arg, '@') {
		return types.NewJID(arg, types.DefaultUserServer), true
	} else {
		recipient, err := types.ParseJID(arg)
		if err != nil {
			return recipient, false
		} else if recipient.User == "" {
			return recipient, false
		}
		return recipient, true
	}
}

func (client *NewClientImpl) FetchGroupAdmin(Jid types.JID) ([]string, error) {
	var Admin []string
	resp, err := client.WA.GetGroupInfo(Jid)
	if err != nil {
		return Admin, err
	} else {
		for _, group := range resp.Participants {
			if group.IsAdmin || group.IsSuperAdmin {
				Admin = append(Admin, group.JID.String())
			}
		}
	}
	return Admin, err
}

func (client *NewClientImpl) SendSticker(jid types.JID, data []byte, opts *waProto.ContextInfo) {
	uploaded, err := client.WA.Upload(context.Background(), data, whatsmeow.MediaImage)
	if err != nil {
		fmt.Printf("Failed to upload file: %v\n", err)
	}

	client.WA.SendMessage(context.Background(), jid, &waProto.Message{
		StickerMessage: &waProto.StickerMessage{
			Url:           proto.String(uploaded.URL),
			DirectPath:    proto.String(uploaded.DirectPath),
			MediaKey:      uploaded.MediaKey,
			Mimetype:      proto.String(http.DetectContentType(data)),
			FileEncSha256: uploaded.FileEncSHA256,
			FileSha256:    uploaded.FileSHA256,
			FileLength:    proto.Uint64(uint64(len(data))),
			ContextInfo:   opts,
		},
	})
}

func (client *NewClientImpl) GetBytes(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return []byte{}, err
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return bytes, nil
}
