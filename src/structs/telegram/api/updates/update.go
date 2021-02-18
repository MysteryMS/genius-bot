package updates

import "mystery.tech/m/v2/src/structs/telegram/api"

type Update struct {
	UpdateId      int                `json:"update_id"`
	InlineQuery   *api.InlineQuery   `json:"inline_query"`
	CallbackQuery *api.CallbackQuery `json:"callback_query"`
}
