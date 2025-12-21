package auth

type TokenPair struct {
	AccessToken  string
	RefreshToken string
}

type TokenService interface {
	Generate(userID int64) (*TokenPair, error)
	ValidateAccessToken(token string) (int64, error)
}
