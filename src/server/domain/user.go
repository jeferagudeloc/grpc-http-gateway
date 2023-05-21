package domain

type (
	UserRepository interface {
		SaveUser(UserData) (UserData, error)
	}

	UserData struct {
		Email    string
		Name     string
		LastName string
	}
)
