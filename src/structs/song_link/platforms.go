package song_link

type Platforms struct {
	AppleMusic   Platform `json:"appleMusic"`
	Spotify      Platform `json:"spotify"`
	Youtube      Platform `json:"youtube"`
	YoutubeMusic Platform `json:"youtubeMusic"`
	Deezer       Platform `json:"deezer"`
	AmazonMusic  Platform `json:"amazonMusic"`
}
