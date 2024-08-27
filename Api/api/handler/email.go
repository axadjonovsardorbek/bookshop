package handler

import (
	"context"
	"fmt"
	"time"

	pb "api/genproto/auth"

	"api/api/helper"

	"github.com/go-redis/redis"
	"golang.org/x/exp/rand"
)


func (h *Handler) UserEmailSending(req *pb.EmailRequest) (*pb.InfoResponse, error) {
	code := rand.Intn(899999) + 10000

	from := "muhammadjonxudaynazarov226@gmail.com"
	password := "rgkt oivo pyst zplt"
	err := helper.SendVerificationCode(helper.Params{
		From:     from,
		Password: password,
		To:       req.Email,
		Message:  fmt.Sprintf("Hi, here is your verification code: %d", code),
		Code:     fmt.Sprint(code),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to send verification email: %v", err)
	}

	err = h.Rdb.Set(context.Background(), fmt.Sprint(code), req.Email, time.Minute*10).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to store verification code in Redis: %v", err)
	}

	return &pb.InfoResponse{
		Message: "A verification code has been sent to your email address.",
		Success: true,
	}, nil
}

func (h *Handler) CheckTheVerificationCode(ctx context.Context, req *pb.ResetPasswordRequest) (*pb.InfoResponse, error) {
	// Retrieve the stored email associated with the verification code from Redis
	_, err := h.Rdb.Get(ctx, req.VerificationCode).Result()
	if err == redis.Nil {
		return &pb.InfoResponse{
			Message: "Verification code is invalid or expired",
			Success: false,
		}, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve verification code from Redis: %v", err)
	}

	err = h.Rdb.Del(ctx, req.VerificationCode).Err()
	if err != nil {
		return nil, fmt.Errorf("failed to delete verification code from Redis: %v", err)
	}

	return &pb.InfoResponse{
		Message: "Verification code is valid and password has been updated",
		Success: true,
	}, nil
}
