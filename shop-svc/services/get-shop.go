package services

import (
	"context"
	"os"
	"time"

	// "log"
	"net/http"

	pb "github.com/asadlive84/bizspace/proto/shop/pb"
	"google.golang.org/protobuf/types/known/timestamppb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) GetShop(ctx context.Context, req *pb.GetShopRequest) (*pb.GetShopResponse, error) {
	var logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"shop_id": req.Id}).Info("Service: GetShop method")

	shop, err := s.Q.GetShopByID(ctx, req.Id)
	if err != nil {
		logger.WithFields(log.Fields{"shop_id": req.Id}).Errorf("an error in get shop query %+v", err)
		return &pb.GetShopResponse{
			Status:  http.StatusNotFound,
			Error:   err.Error(),
			Message: "shop not found",
		}, nil
	}

	logger.WithFields(log.Fields{"shop_id": req.Id}).Info("Successfully retrieved shop")

	shop1 := &pb.Shop{
		// Id:    shop,
		Name:      shop.Name,
		Address:   shop.Address,
		OwnerId:   shop.OwnerID,
		CreatedAt: timestamppb.New(shop.CreatedAt),
		UpdatedAt: timestamppb.New(shop.UpdatedAt.Time),
	}

	return &pb.GetShopResponse{
		Shop:    shop1,
		Status:  http.StatusOK,
		Message: "Shop retrieved successfully",
	}, nil
}
