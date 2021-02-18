package api

type AnswerInline struct {
	InlineQueryId string                     `json:"inline_query_id"`
	Results       []InlineQueryResultArticle `json:"results"`
}
