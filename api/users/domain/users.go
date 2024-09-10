package users_domain

type Users struct {
	ID       string
	Username string `validate:"required"`
	Password string `validate:"required"`
}
