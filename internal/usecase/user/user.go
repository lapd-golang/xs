package user

import (
	"errors"

	"github.com/zainul/xs/internal/entity"
	"github.com/zainul/xs/internal/pkg/error/deliveryerror"
	"github.com/zainul/xs/internal/pkg/error/usecaseerror"
	"github.com/zainul/xs/internal/repository"
	"github.com/zainul/xs/internal/usecase"
)

type userUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase ...
func NewUserUseCase(userRepo repository.UserRepository) usecase.User {
	return &userUseCase{
		userRepo: userRepo,
	}
}

func (u *userUseCase) Register(user entity.User) *deliveryerror.Error {
	// validation must be in here
	err := user.Validate()

	if err != nil {
		return deliveryerror.GetError(usecaseerror.ValidationFailed, err)
	}

	user.AccountNumber, err = u.userRepo.GenerateAccountNumber()

	if err != nil {
		return deliveryerror.GetError(usecaseerror.FailedToGenerateAccountNumber, err)
	}

	errRepo := u.userRepo.Save(user)

	return deliveryerror.GetError(usecaseerror.InternalServerError, errRepo)
}

func (u *userUseCase) EditProfile(user entity.User) *deliveryerror.Error {
	// validation must be in here

	if userInfo := u.userRepo.GetByField(user.Email, "email"); userInfo != nil {
		errRepo := u.userRepo.Edit(user, user.Email)
		return deliveryerror.GetError(usecaseerror.InternalServerError, errRepo)
	}

	return deliveryerror.GetError(
		usecaseerror.UserNotFound,
		errors.New(usecaseerror.UserNotFound),
	)
}

func (u *userUseCase) ResetPassword(email string) *deliveryerror.Error {
	return nil
}

func (u *userUseCase) DefaultInfo(accountNumber string) *entity.User {
	return u.userRepo.GetByField(accountNumber, "eamil")
}
