package api

type CallbackQuery struct {
	Id              string `json:"id"`
	Message         string `json:"message"`
	InlineMessageId string `json:"inline_message_id"`
	Data            string `json:"data"`
}
