package typings

type MediaType string

const (
	IMAGE = "image"
	VIDEO = "video"
)

type MetadataSticker struct {
	Author    string
	Pack      string
	KeepScale bool
	Removebg  any
	Circle    bool
}

type Sticker struct {
	File []byte
	Tipe MediaType
}
