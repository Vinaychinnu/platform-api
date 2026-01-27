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

type CreateReleaseRequest struct {
	ID            string `json:"id"`
	EnvironmentID string `json:"environment_id"`
	Version       string `json:"version"`
}

type CreateReleaseResponse struct {
	ID            string `json:"id"`
	EnvironmentID string `json:"environment_id"`
	Version       string `json:"version"`
	Status        string `json:"status"`
}
