package request

type AddTaskRequest struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
	Username    string `json:"assignedUser"`
}
