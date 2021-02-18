package updates

import "mystery.tech/m/v2/src/structs/telegram/api"

type MarkupMessage struct {
	InlineMessageId string                   `json:"inline_message_id"`
	ReplyMarkup     api.InlineKeyboardMarkup `json:"reply_markup"`
}
