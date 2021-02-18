package webhook

type Webhook struct {
	Url     string `json:"url"`
	Pending int64  `json:"pending_update_count"`
	Error   string `json:"last_error_message"`
}
