package param

// CreateRepositoryRequest
type CreateRepositoryRequest struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// CreateRepositoryResponse
type CreateRepositoryResponse struct {
	ID int `json:"id"`
}
