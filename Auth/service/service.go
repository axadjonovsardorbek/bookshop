package service

import (
	"context"
	"log"

	pb "auth/genproto/auth"
	stg "auth/storage"
)

type AuthService struct {
	stg stg.InitRoot
	pb.UnimplementedAuthServiceServer
}

func NewAuthService(stg stg.InitRoot) *AuthService {
	return &AuthService{
		stg: stg,
	}
}

func (s *AuthService) SignUp(ctx context.Context, req *pb.Users) (*pb.SignUpResponse, error) {
	resp, err := s.stg.Auth().SignUp(req)
	if err != nil {
		log.Println("Error signing up: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) LogIn(ctx context.Context, req *pb.LogInRequest) (*pb.LogInResponse, error) {
	resp, err := s.stg.Auth().LogIn(req)
	if err != nil {
		log.Println("Error logging in: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ChangePassword(ctx context.Context, req *pb.ChangePasswordRequest) (*pb.InfoResponse, error) {
	resp, err := s.stg.Auth().ChangePassword(req)
	if err != nil {
		log.Println("Error changing password: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ForgetPassword(ctx context.Context, req *pb.ForgetPasswordRequest) (*pb.InfoResponse, error) {
	resp, err := s.stg.Auth().ForgetPassword(req)
	if err != nil {
		log.Println("Error handling forget password: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ResetPassword(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.InfoResponse, error) {
	resp, err := s.stg.Auth().ResetPassword(req)
	if err != nil {
		log.Println("Error resetting password: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ChangeEmail(ctx context.Context, req *pb.ChangeEmailRequest) (*pb.InfoResponse, error) {
	resp, err := s.stg.Auth().ChangeEmail(req)
	if err != nil {
		log.Println("Error changing email: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) VerifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.InfoResponse, error) {
	resp, err := s.stg.Auth().VerifyEmail(req)
	if err != nil {
		log.Println("Error verifying email: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) EnterEmail(ctx context.Context, req *pb.EmailRequest) (*pb.InfoResponse, error) {
	resp, err := s.stg.Auth().EnterEmail(req)
	if err != nil {
		log.Println("Error entering email: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateTokenRequest) (*pb.InfoResponse, error) {
	resp, err := s.stg.Auth().ValidateToken(req)
	if err != nil {
		log.Println("Error validating token: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) RefreshToken(ctx context.Context, req *pb.RefreshTokenRequest) (*pb.InfoResponse, error) {
	resp, err := s.stg.Auth().RefreshToken(req)
	if err != nil {
		log.Println("Error refreshing token: ", err)
		return nil, err
	}
	return resp, nil
}

func (s *AuthService) GetProfile(ctx context.Context, req *pb.GetProfileRequest) (*pb.GetProfileResponse, error) {
	resp, err := s.stg.Auth().GetProfile(req)
	if err != nil {
		log.Println("Error while getting the profile: ", err)
		return nil, err
	}
	return resp, nil
}
