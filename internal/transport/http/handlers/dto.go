package handlers

type CreateProjectRequest struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateProjectResponse struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

type CreateEnvironmentRequest struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
}

type CreateEnvironmentResponse struct {
	ID        string `json:"id"`
	ProjectID string `json:"project_id"`
	Name      string `json:"name"`
}
