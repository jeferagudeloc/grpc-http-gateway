package domain

type (
	UserRepository interface {
		GetUsers() ([]User, error)
	}

	User struct {
		Name     string `json:"name"`
		LastName string `json:"lastname"`
		Email    string `json:"email"`
		Status   string `json:"status"`
		Role     Role   `json:"role"`
	}

	Role struct {
		Name        string   `json:"name"`
		Permissions []string `json:"permissions"`
	}
)
