package handlers

type CreateProjectRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateProjectResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
