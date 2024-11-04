package usecases

type AuthServiceInterface interface {
	GenerateToken(login string, password string) (string, error)
	ParseToken(token string) (int, error)
}
