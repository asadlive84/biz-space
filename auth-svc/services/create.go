package services

import (
	"context"
	"database/sql"
	"os"
	"time"

	// "log"
	"net/http"

	q "github.com/asadlive84/bizspace/auth-svc/internal/query"
	pb "github.com/asadlive84/bizspace/proto/auth/pb"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/timestamppb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"email": req.Email}).Info("Service: Register method")

	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(req.PasswordHash), 14)

	if err != nil {
		logger.WithFields(log.Fields{"email": req.Email}).Errorf("an error in hashed password %+v", err)
		return &pb.CreateUserResponse{
			Status: http.StatusBadRequest,
			Error:  "account not created",
		}, nil
	}

	user, err := s.Q.GetUserByEmail(req.Email)
	if err != nil && err != q.NotFound {
		logger.WithFields(log.Fields{"email": req.Email}).Errorf("an error in get user query %+v", err)
		return &pb.CreateUserResponse{
			Status:  http.StatusBadRequest,
			Error:   "query issue",
			Message: "an error in get user query",
		}, nil
	}

	if user != nil {
		logger.WithFields(log.Fields{"email": req.Email}).Info("email is exists")
		return &pb.CreateUserResponse{
			Status:  http.StatusBadRequest,
			Error:   "email is exits",
			Message: "this user email address already exists in our system",
		}, nil
	}

	// if req.GetPhoneNumber() != "" {
	// 	user, err = s.Q.GetUserByPhone(req.GetPhoneNumber())
	// 	if err != nil && err != q.NotFound {
	// 		logger.WithFields(log.Fields{"phone": req.Email}).Errorf("an error in get user query %+v", err)
	// 		return &pb.CreateUserResponse{
	// 			Status: http.StatusBadRequest,
	// 		}, nil
	// 	}

	// 	if user != nil {
	// 		logger.WithFields(log.Fields{"phone": req.Email}).Info("phone is exists")
	// 		return &pb.CreateUserResponse{
	// 			Status: http.StatusBadRequest,
	// 		}, nil
	// 	}
	// }

	result, err := s.Q.CreateUser(q.User{
		UserName:     req.GetUserName(),
		FullName:     req.GetFullName(),
		PhoneNumber:  req.GetPhoneNumber(),
		Address:      req.GetAddress(),
		PasswordHash: string(hashedPasswd),
		Email:        req.GetEmail(),
	})

	if err != nil {
		logger.WithFields(log.Fields{"email": req.Email}).Errorf("an error in query %+v", err)
		return &pb.CreateUserResponse{
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
			Message: "an error in query",
		}, nil
	}

	logger.WithFields(log.Fields{"email": req.Email}).Info("Successfully created user")

	user1 := &pb.User{
		UserId:      result.UserID,
		UserName:    result.UserName,
		FullName:    result.FullName,
		PhoneNumber: result.PhoneNumber,
		Address:     result.Address,
		Email:       result.Email,
		CreatedAt:   timestamppb.New(result.CreatedAt),
	}

	return &pb.CreateUserResponse{
		User:    user1,
		Status:  http.StatusCreated,
		Message: "User created success",
	}, nil
}

func createNullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{
			String: "",
			Valid:  false,
		}
	}
	return sql.NullString{
		String: s,
		Valid:  true,
	}
}
