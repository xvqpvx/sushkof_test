package model

type Task struct {
	IdTask      int
	Title       string
	Description string
	Status      string
	IsActive    bool
	CreatedAt   string
	UpdatedAt   string
	UserId      int
}
