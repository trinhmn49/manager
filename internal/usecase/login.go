package usecase

import (
	"context"
	"errors"
	"manager/internal/domain/repository"
	"manager/pkg/hash"
	"manager/pkg/validator"

	"gorm.io/gorm"
)

type LoginDTO struct {
	ID          string  `json:"id"`
	DisplayName string  `json:"display_name"`
	Avatar      string  `json:"avatar"`
	Phone       string  `json:"phone"`
	Email       *string `json:"email"`

	Token TokenRes `json:"tokens"`
}

type TokenRes struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type LoginInput struct {
	Phone    string
	Password string
}

type LoginUseCase interface {
	Execute(ctx context.Context, input LoginInput) (*LoginDTO, error)
}

type loginUseCaseImpl struct {
	userRepo repository.UserRepo
	hasher   hash.PasswordHasher
}

func (l *loginUseCaseImpl) Execute(ctx context.Context, input LoginInput) (*LoginDTO, error) {
	if !validator.IsPhoneValid(input.Phone) {
		return nil, errors.New("phone invalid")
	}

	// Find user info with phone
	user, err := l.userRepo.FindUserByPhone(ctx, input.Phone)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, gorm.ErrRecordNotFound
		}
		return nil, err
	}

	// verify password
	err = l.hasher.Verify(user.Password, input.Password)
	if err != nil {
		return nil, errors.New("err verify password")
	}

	// handle response
	loginDTO := LoginDTO{
		//ID:          user.ID.String(),
		DisplayName: user.DisplayName,
		Avatar:      user.Avatar,
		Phone:       user.Phone,
		Email:       user.Email,
	}

	//Create token for register session - handle later
	/*
		accessTokenExpTime := time.Now().Add(time.Hour * 24)
		refreshTokenExpTime := time.Now().Add(time.Hour * 24 * 7)

		accessToken, err := token.GenToken(loginDTO.ID, accessTokenExpTime)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("internal server errors")
		}

		refreshToken, err := token.GenToken(loginDTO.ID, refreshTokenExpTime)
		if err != nil {
			fmt.Println(err)
			return nil, fmt.Errorf("internal server errors")
		}

		loginDTO.Token = TokenRes{
			AccessToken:  accessToken,
			RefreshToken: refreshToken,
		}
	*/

	return &loginDTO, nil
}

func NewLoginUseCase(userRepo repository.UserRepo, hasher hash.PasswordHasher) LoginUseCase {
	return &loginUseCaseImpl{
		userRepo: userRepo,
		hasher:   hasher,
	}
}
