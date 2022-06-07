package service

import (
	"context"
	"github.com/samithiwat/samithiwat-backend-gateway/src/dto"
	"github.com/samithiwat/samithiwat-backend-gateway/src/proto"
	"log"
	"net/http"
	"time"
)

type AuthService struct {
	client proto.AuthServiceClient
}

func NewAuthService(client proto.AuthServiceClient) *AuthService {
	return &AuthService{
		client: client,
	}
}

func (s *AuthService) Register(register *dto.Register) (result *proto.User, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	r := s.DtoToRawRegister(register)

	res, errRes := s.client.Register(ctx, &proto.RegisterRequest{Register: r})

	if errRes != nil {
		log.Printf("%v\n", errRes)
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusCreated {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data
	return
}

func (s *AuthService) Login(login *dto.Login) (result *proto.Credential, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	l := s.DtoToRawLogin(login)

	res, errRes := s.client.Login(ctx, &proto.LoginRequest{Login: l})
	if errRes != nil {
		log.Printf("%v\n", errRes)
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data
	return
}

func (s *AuthService) ChangePassword(changePwd *dto.ChangePassword) (result bool, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	chPwd := s.DtoToRawChangePassword(changePwd)

	res, errRes := s.client.ChangePassword(ctx, &proto.ChangePasswordRequest{ChangePassword: chPwd})
	if errRes != nil {
		log.Printf("%v\n", errRes)
		return false, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusNoContent {
		return false, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data
	return
}

func (s *AuthService) Logout(userId uint32) (result bool, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.Logout(ctx, &proto.LogoutRequest{UserId: uint32(userId)})
	if errRes != nil {
		log.Printf("%v\n", errRes)
		return false, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusNoContent {
		return false, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data
	return
}

func (s *AuthService) Validate(token string) (result uint32, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.Validate(ctx, &proto.ValidateRequest{Token: token})
	if errRes != nil {
		log.Println(errRes)
		return 0, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return 0, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data
	return
}

func (s *AuthService) RefreshToken(token string) (result *proto.Credential, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	res, errRes := s.client.RefreshToken(ctx, &proto.RefreshTokenRequest{RefreshToken: token})
	if errRes != nil {
		log.Println(errRes)
		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	if res.StatusCode != http.StatusOK {
		return nil, &dto.ResponseErr{
			StatusCode: int(res.StatusCode),
			Message:    FormatErr(res.Errors),
			Data:       nil,
		}
	}

	result = res.Data
	return
}

func (s *AuthService) DtoToRawRegister(register *dto.Register) *proto.Register {
	return &proto.Register{
		Email:       register.Email,
		Password:    register.Password,
		Firstname:   register.Firstname,
		DisplayName: register.DisplayName,
		Lastname:    register.Lastname,
		ImageUrl:    register.ImageUrl,
	}
}

func (s *AuthService) DtoToRawLogin(login *dto.Login) *proto.Login {
	return &proto.Login{
		Email:    login.Email,
		Password: login.Password,
	}
}

func (s *AuthService) DtoToRawChangePassword(chPwd *dto.ChangePassword) *proto.ChangePassword {
	return &proto.ChangePassword{
		UserId:      chPwd.UserId,
		OldPassword: chPwd.OldPassword,
		NewPassword: chPwd.NewPassword,
	}
}
