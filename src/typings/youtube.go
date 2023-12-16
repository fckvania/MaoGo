package typings

type YoutubeInfo struct {
	Title          string
	Duration       float64
	ParsedDuration string
	Author         string
}

type YoutubeAV struct {
	Size    string
	Format  string
	Quality string
	Url     func() string
}

type YoutubeLinks struct {
	Audio []YoutubeAV
	Video []YoutubeAV
}

type YoutubeInfos struct {
	Info YoutubeInfo
	Link YoutubeLinks
}
