package response

type CategoryResponse struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	ID          int    `json:"id"`
}
