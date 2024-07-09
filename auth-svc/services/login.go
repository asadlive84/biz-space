package services

import (
	"context"
	"os"
	"time"

	// "log"
	"net/http"

	q "github.com/asadlive84/bizspace/auth-svc/internal/query"
	pb "github.com/asadlive84/bizspace/proto/auth/pb"
	"golang.org/x/crypto/bcrypt"

	log "github.com/sirupsen/logrus"
)

func (s *Server) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {

	var logger = log.New()

	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	},
	)

	// Set the output to standard output (console)
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"email": req.Email}).Info("Service: Login method with email")
	logger.WithFields(log.Fields{"phone": req.Phone}).Info("Service: Login method with phone")

	user, err := s.Q.GetUserByEmail(req.Email)
	if err != nil && err != q.NotFound {
		logger.WithFields(log.Fields{"email": req.Email}).Errorf("an error in get user query %+v", err)
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	if user == nil {
		user, err = s.Q.GetUserByPhone(req.Phone)
		if err != nil && err != q.NotFound {
			logger.WithFields(log.Fields{"phone": req.Phone}).Errorf("an error in get user query %+v", err)
			return &pb.LoginResponse{
				Status: http.StatusBadRequest,
			}, nil
		}
	}

	if user == nil {
		logger.Errorln("user email or phone is not found")
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, nil
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password))
	if err != nil {
		logger.WithFields(log.Fields{"email": req.Email}).Errorf("has mismatched %+v", err)
		return &pb.LoginResponse{
			Status: http.StatusBadRequest,
		}, nil
	}
	logger.WithFields(log.Fields{"email": req.Email}).Info("login successfully")
	return &pb.LoginResponse{
		Status: http.StatusCreated,
	}, nil
}
