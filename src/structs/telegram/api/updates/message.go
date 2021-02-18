package updates

import "mystery.tech/m/v2/src/structs/telegram/api"

type Message struct {
	InlineMessageId       string                   `json:"inline_message_id"`
	Text                  string                   `json:"text"`
	ParseMode             string                   `json:"parse_mode"`
	ReplyMarkup           api.InlineKeyboardMarkup `json:"reply_markup"`
	DisableWebpagePreview bool                     `json:"disable_webpage_preview"`
}
