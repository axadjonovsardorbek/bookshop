package storage

import (
	"auth/genproto/auth"
)

type InitRoot interface {
	Auth() AuthService
}

type AuthService interface {
	SignUp(req *auth.Users) (*auth.SignUpResponse, error)
	LogIn(req *auth.LogInRequest) (*auth.LogInResponse, error)
	ChangePassword(req *auth.ChangePasswordRequest) (*auth.InfoResponse, error)
	ForgetPassword(req *auth.ForgetPasswordRequest) (*auth.InfoResponse, error)
	ResetPassword(req *auth.ResetPasswordRequest) (*auth.InfoResponse, error)
	ChangeEmail(req *auth.ChangeEmailRequest) (*auth.InfoResponse, error)
	VerifyEmail(req *auth.VerifyEmailRequest) (*auth.InfoResponse, error)
	EnterEmail(req *auth.EmailRequest) (*auth.InfoResponse, error)
	ValidateToken(req *auth.ValidateTokenRequest) (*auth.InfoResponse, error)
	RefreshToken(req *auth.RefreshTokenRequest) (*auth.InfoResponse, error)
	GetProfile(req *auth.GetProfileRequest) (*auth.GetProfileResponse, error)
}
