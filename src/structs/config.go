package structs

type Config struct {
	Telegram string `json:"telegramToken"`
	Genius   string `json:"geniusToken"`
	SongLink string `json:"songLinkToken"`
	Webhook  string `json:"webhookUrl"`
	Debug    bool   `json:"debug"`
}
