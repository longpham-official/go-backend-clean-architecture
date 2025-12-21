package auth

import (
	domainAuth  "github.com/longpham-official/go-backend-clean-architecture/internal/domain/auth"
	domainUser "github.com/longpham-official/go-backend-clean-architecture/internal/domain/user"
	"errors"
	"context"
)

var (
	ErrInvalidCredential = errors.New("invalid email or password")
)

type LoginInput struct {
	Email    string
	Password string
}

type LoginOutput struct {
	AccessToken string
	RefreshToken string
}

type LoginUsecase struct {
	userRepo	domainUser.Repository
	tokenSvc	domainAuth.TokenService
	passwordFn func(hashed, plain  string) bool
}

func NewLoginUsecase(
	userRepo domainUser.Repository,
	tokenSvc domainAuth.TokenService,
	passwordFn func(hashed, plain string) bool,
) *LoginUsecase {
	return &LoginUsecase{
		userRepo:    userRepo,
		tokenSvc:   tokenSvc,
		passwordFn: passwordFn,
	}
}

func (uc *LoginUsecase) Execute(
	ctx context.Context,
	input LoginInput,
) (*LoginOutput, error) {

	user, err := uc.userRepo.GetUserByEmail(ctx, input.Email)
	if err != nil || user == nil {
		return nil, ErrInvalidCredential
	}

	if !uc.passwordFn(user.Password, input.Password) {
		return nil, ErrInvalidCredential
	}

	tokenPair, err := uc.tokenSvc.Generate(user.ID)
	if err != nil {
		return nil, err
	}

	return &LoginOutput{
		AccessToken:  tokenPair.AccessToken,
		RefreshToken: tokenPair.RefreshToken,
	}, nil
}
