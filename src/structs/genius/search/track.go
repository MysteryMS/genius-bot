package search

type Track struct {
	Id    int    `json:"id"`
	Title string `json:"full_title"`
	Url   string `json:"url"`
	Stats Stats  `json:"stats"`
	Cover string `json:"song_art_image_thumbnail_url"`
}
