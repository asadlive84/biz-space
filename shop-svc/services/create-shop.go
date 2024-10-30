package services

import (
	"context"
	"os"
	"time"

	// "log"
	"net/http"

	pb "github.com/asadlive84/bizspace/proto/shop/pb"
	q "github.com/asadlive84/bizspace/shop-svc/internal/query"
	"google.golang.org/protobuf/types/known/timestamppb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) CreateShop(ctx context.Context, req *pb.CreateShopRequest) (*pb.CreateShopResponse, error) {
	var logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"name": req.Name}).Info("Service: CreateShop method")

	shop, err := s.Q.GetShopByName(ctx, req.Name)
	if err != nil && err != q.NotFound {
		logger.WithFields(log.Fields{"name": req.Name}).Errorf("an error in get shop query %+v", err)
		return &pb.CreateShopResponse{
			Status:  http.StatusBadRequest,
			Error:   "query issue",
			Message: "1 an error in get shop query",
		}, nil
	}

	if (shop != q.Shop{}) {
		logger.WithFields(log.Fields{"name": req.Name}).Info("shop name exists")
		return &pb.CreateShopResponse{
			Status:  http.StatusBadRequest,
			Error:   "shop name exists",
			Message: "2 this shop name already exists in our system",
		}, nil
	}

	result, err := s.Q.CreateShop(ctx, q.Shop{
		Name:    req.GetName(),
		Address: req.GetAddress(),
		OwnerID: req.GetOwnerId(),
	})

	if err != nil {
		logger.WithFields(log.Fields{"name": req.Name}).Errorf("an error in query %+v", err)
		return &pb.CreateShopResponse{
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
			Message: "3 an error in query",
		}, nil
	}

	logger.WithFields(log.Fields{"name": req.Name}).Info("Successfully created shop")

	shop1 := &pb.Shop{
		Id:        result.ID,
		Name:      result.Name,
		Address:   result.Address,
		OwnerId:   result.OwnerID,
		CreatedAt: timestamppb.New(result.CreatedAt),
	}

	return &pb.CreateShopResponse{
		Shop:    shop1,
		Status:  http.StatusCreated,
		Message: "Shop created successfully",
	}, nil
}
