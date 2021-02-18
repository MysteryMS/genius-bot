package search

type Response struct {
	Hits *[]Hit `json:"hits"`
}
