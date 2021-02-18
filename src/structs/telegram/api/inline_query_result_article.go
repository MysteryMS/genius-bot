package api

type InlineQueryResultArticle struct {
	Type                string                  `json:"type"`
	Id                  string                  `json:"id"`
	Title               string                  `json:"title"`
	InputMessageContent InputTextMessageContent `json:"input_message_content"`
	ReplyMarkup         InlineKeyboardMarkup    `json:"reply_markup"`
	ThumbUrl            string                  `json:"thumb_url"`
}
