package services

import (
	"context"
	"crypto/rand"
	"database/sql"
	"encoding/base64"
	"errors"
	"net/http"
	"time"

	pb "github.com/asadlive84/bizspace/proto/auth/pb"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	q "github.com/asadlive84/bizspace/auth-svc/internal/query"
)

// Generate a random refresh token
func generateRefreshToken() (string, error) {
	b := make([]byte, 32)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(b), nil
}

// Generate a JWT token
func (s *Server) generateToken(user *q.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.UserID,
		"exp":     time.Now().Add(time.Hour * 3).Unix(), // Token expiry time
	})

	return token.SignedString([]byte(s.jwtSecret))
}

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	var user *q.User
	var err error

	if req.Email != "" {
		user, err = s.Q.GetUserByEmail(req.Email)
	} else if req.Phone != "" {
		user, err = s.Q.GetUserByPhone(req.Phone)
	} else {
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
			Error:  "Email or phone is required",
		}, nil
	}

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return &pb.LoginResponse{
				Status: http.StatusUnauthorized,
				Error:  "Invalid credentials",
			}, nil
		}
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "Database error",
		}, nil
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return &pb.LoginResponse{
			Status: http.StatusUnauthorized,
			Error:  "Invalid credentials",
		}, nil
	}

	accessToken, err := s.generateToken(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to generate token",
		}, nil
	}

	refreshToken, err := generateRefreshToken()
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to generate refresh token",
		}, nil
	}

	user.RefreshToken.String = refreshToken
	err = s.Q.UpdateUserRefreshToken(user)
	if err != nil {
		return &pb.LoginResponse{
			Status: http.StatusInternalServerError,
			Error:  "Failed to store refresh token",
		}, nil
	}

	return &pb.LoginResponse{
		Status:       http.StatusOK,
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}
