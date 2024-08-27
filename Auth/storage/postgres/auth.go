package postgres

import (
	"auth/genproto/auth"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type AuthStorage struct {
	db *sql.DB
}

func NewAuthStorage(db *sql.DB) *AuthStorage {
	return &AuthStorage{
		db: db,
	}
}

func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

func (s *AuthStorage) SignUp(req *auth.Users) (*auth.SignUpResponse, error) {

	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		return nil, fmt.Errorf("failed to hash password: %v", err)
	}

	userQuery := `
		INSERT INTO users (
			id, first_name, last_name, email, password, date_of_birth, phone_number
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7
		)
		RETURNING id
	`

	var userId string

	err = s.db.QueryRowContext(
		context.TODO(),
		userQuery,
		req.UserId, req.FirstName, req.LastName, req.Email, hashedPassword, req.DateOfBirth, req.PhoneNumber,
	).Scan(&userId)

	if err != nil {
		return nil, fmt.Errorf("failed to insert user: %v", err)
	}

	accessToken := req.AccessToken   // Or your token generation logic
	refreshToken := req.RefreshToken // Or your token generation logic
	expiresAt := time.Now().Add(24 * time.Hour).Format(time.RFC3339)

	// SQL query to insert tokens
	tokenQuery := `
		INSERT INTO tokens (
			token_id, user_id, access_token, refresh_token, expires_at
		) VALUES (
			$1, $2, $3, $4, $5
		)
	`

	_, err = s.db.ExecContext(
		context.TODO(),
		tokenQuery,
		uuid.NewString(), userId, accessToken, refreshToken, expiresAt,
	)

	if err != nil {
		return nil, fmt.Errorf("failed to insert tokens: %v", err)
	}

	response := &auth.SignUpResponse{
		UserId:       userId,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	return response, nil
}

func (u *AuthStorage) LogIn(req *auth.LogInRequest) (*auth.LogInResponse, error) {
	var storedPasswordHash string
	var userID string

	query := `
		SELECT id, password
		FROM users
		WHERE email = $1
	`
	err := u.db.QueryRow(query, req.Email).Scan(&userID, &storedPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(req.Password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	tokenQuery := `
		SELECT access_token, refresh_token, expires_at
		FROM tokens
		WHERE user_id = $1
	`
	var accessToken, refreshToken string
	var expiresAt sql.NullString

	err = u.db.QueryRow(tokenQuery, userID).Scan(&accessToken, &refreshToken, &expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("no tokens found for the user")
		}
		return nil, err
	}

	// Handle the case where expires_at is NULL
	// expiresAtStr := ""
	// if expiresAt.Valid {
	// 	expiresAtStr = expiresAt.String
	// }

	response := &auth.LogInResponse{
		UserId:       userID,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
		// ExpiresAt:    expiresAtStr,
	}

	return response, nil
}

func (u *AuthStorage) ChangePassword(req *auth.ChangePasswordRequest) (*auth.InfoResponse, error) {
	var currentPasswordHash string
	query := `
		SELECT password
		FROM users
		WHERE id = $1 AND deleted_at = 0
	`
	err := u.db.QueryRow(query, req.UserId).Scan(&currentPasswordHash)
	if err != nil {
		if err == sql.ErrNoRows {
			return &auth.InfoResponse{Success: false, Message: "User not found or already deleted"}, nil
		}
		return nil, fmt.Errorf("could not fetch user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(currentPasswordHash), []byte(req.OldPassword))
	if err != nil {
		return &auth.InfoResponse{Success: false, Message: "Current password is incorrect"}, nil
	}

	hashedNewPassword, err := hashPassword(req.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("could not hash new password: %v", err)
	}

	updateQuery := `
		UPDATE users
		SET password = $1, updated_at = CURRENT_TIMESTAMP
		WHERE id = $2 AND deleted_at = 0
	`

	result, err := u.db.Exec(updateQuery, hashedNewPassword, req.UserId)
	if err != nil {
		return &auth.InfoResponse{Success: false, Message: "Failed to update password"}, fmt.Errorf("could not change password: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &auth.InfoResponse{Success: false, Message: "Failed to update password"}, fmt.Errorf("could not determine rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return &auth.InfoResponse{Success: false, Message: "Failed to update password"}, fmt.Errorf("no rows were updated")
	}

	return &auth.InfoResponse{Success: true, Message: "Password successfully changed"}, nil
}

func (s *AuthStorage) ResetPassword(req *auth.ResetPasswordRequest) (*auth.InfoResponse, error) {
	hashedNewPassword, err := hashPassword(req.NewPassword)
	if err != nil {
		return nil, fmt.Errorf("could not hash new password: %v", err)
	}

	updateQuery := `
		UPDATE users
		SET password = $1, updated_at = CURRENT_TIMESTAMP
		WHERE email = $2 AND deleted_at = 0
		`
	result, err := s.db.Exec(updateQuery, hashedNewPassword, req.Email)
	if err != nil {
		return &auth.InfoResponse{Success: false, Message: "Failed to reset password"}, fmt.Errorf("could not reset password: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return &auth.InfoResponse{Success: false, Message: "Failed to reset password"}, fmt.Errorf("could not determine rows affected: %v", err)
	}
	if rowsAffected == 0 {
		return &auth.InfoResponse{Success: false, Message: "Failed to reset password"}, fmt.Errorf("no rows were updated")
	}

	return &auth.InfoResponse{Success: true, Message: "Password successfully reset"}, nil
}

func (s *AuthStorage) ChangeEmail(req *auth.ChangeEmailRequest) (*auth.InfoResponse, error) {
	tx, err := s.db.BeginTx(context.Background(), nil)
	if err != nil {
		return nil, fmt.Errorf("failed to start transaction: %v", err)
	}
	defer tx.Rollback()

	var storedPasswordHash string
	var userID string
	err = tx.QueryRow(`SELECT id, password FROM users WHERE email = $1`, req.CurrentEmail).Scan(&userID, &storedPasswordHash)
	if err == sql.ErrNoRows {
		return &auth.InfoResponse{
			Message: "Email not found",
			Success: false,
		}, nil
	} else if err != nil {
		return nil, fmt.Errorf("failed to retrieve user: %v", err)
	}

	err = bcrypt.CompareHashAndPassword([]byte(storedPasswordHash), []byte(req.Password))
	if err != nil {
		return &auth.InfoResponse{
			Message: "Invalid password",
			Success: false,
		}, nil
	}

	_, err = tx.Exec(`UPDATE users SET email = $1, is_email_verified = $2 WHERE id = $3`,
		req.NewEmail, false, userID)
	if err != nil {
		return nil, fmt.Errorf("failed to update email: %v", err)
	}

	if err := tx.Commit(); err != nil {
		return nil, fmt.Errorf("failed to commit transaction: %v", err)
	}

	return &auth.InfoResponse{
		Message: "Email updated successfully, please verify your new email address",
		Success: true,
	}, nil
}

func (s *AuthStorage) VerifyEmail(req *auth.VerifyEmailRequest) (*auth.InfoResponse, error) {
	ctx := context.Background()

	query := "UPDATE users SET is_email_verified = true  WHERE email = $1 RETURNING email;"

	var email string
	err := s.db.QueryRowContext(ctx, query, req.Email).Scan(&email)
	if err != nil {
		if err == sql.ErrNoRows {
			return &auth.InfoResponse{
				Message: "Email not found",
				Success: false,
			}, nil
		}
		return nil, fmt.Errorf("failed to update email verification status: %v", err)
	}

	return &auth.InfoResponse{
		Message: "Email verified successfully",
		Success: true,
	}, nil
}

func (u *AuthStorage) ValidateToken(req *auth.ValidateTokenRequest) (*auth.InfoResponse, error) {
	var expiresAt string

	query := `
		SELECT expires_at
		FROM tokens
		WHERE access_token = $1
	`

	err := u.db.QueryRow(query, req.Token).Scan(&expiresAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return &auth.InfoResponse{
				Message: "Token does not exist in the database!",
				Success: false,
			}, nil
		}
		return nil, err
	}

	expiredAt, err := time.Parse(time.RFC3339, expiresAt)
	if err != nil {
		log.Fatalln("Error while parsing the time", err)
	}

	if time.Now().After(expiredAt) {
		return &auth.InfoResponse{
			Message: "Token is already expired",
			Success: false,
		}, nil
	}

	return &auth.InfoResponse{
		Message: "Token is valid",
		Success: true,
	}, nil
}

func (u *AuthStorage) GetProfile(req *auth.GetProfileRequest) (*auth.GetProfileResponse, error) {
	query := `
		SELECT first_name, last_name, email, phone_number, date_of_birth
		FROM users
		WHERE id = $1
	`

	row := u.db.QueryRow(query, req.UserId)

	resp := &auth.GetProfileResponse{}

	err := row.Scan(
		&resp.FirstName,
		&resp.LastName,
		&resp.Email,
		&resp.PhoneNumber,
		&resp.DateOfBirth,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("user not found")
		}
		return nil, fmt.Errorf("error fetching profile: %w", err)
	}

	return resp, nil
}

func (u *AuthStorage) RefreshToken(req *auth.RefreshTokenRequest) (*auth.InfoResponse, error) {
	return nil, nil
}

func (s *AuthStorage) EnterEmail(req *auth.EmailRequest) (*auth.InfoResponse, error) {
	return nil, nil
	// This method is never be used
}

func (s *AuthStorage) ForgetPassword(req *auth.ForgetPasswordRequest) (*auth.InfoResponse, error) {
	// THis method is never be used
	return nil, nil
}
