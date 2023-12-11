package request

type UpdateTaskRequest struct {
	IdTask      int
	Title       string `json:"title,omitempty"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status,omitempty"`
}
