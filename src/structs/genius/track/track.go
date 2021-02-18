package track

import "mystery.tech/m/v2/src/structs/genius/album"

type Track struct {
	AppleMusicId  string        `json:"apple_music_id"`
	Description   Description   `json:"description"`
	Album         *album.Album  `json:"album"`
	ReleaseDate   string        `json:"release_date_for_display"`
	Title         string        `json:"title_with_featured"`
	Cover         string        `json:"song_art_image_url"`
	Url           string        `json:"url"`
	Id            string        `json:"id"`
	PrimaryArtist PrimaryArtist `json:"primary_artist"`
}
