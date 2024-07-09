package services

import (
	"context"
	"database/sql"
	"os"
	"time"

	// "log"
	"net/http"

	pb "github.com/asadlive84/bizspace/proto/shop/pb"
	q "github.com/asadlive84/bizspace/shop-svc/internal/query"
	"google.golang.org/protobuf/types/known/timestamppb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) UpdateShop(ctx context.Context, req *pb.UpdateShopRequest) (*pb.UpdateShopResponse, error) {
	var logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"shop_id": req.Id}).Info("Service: UpdateShop method")

	shop, err := s.Q.UpdateShop(ctx, q.Shop{
		ID:        req.GetId(),
		Name:      req.GetName(),
		Address:   req.GetAddress(),
		OwnerID:   req.GetOwnerId(),
		UpdatedAt: sql.NullTime{Time: time.Now(), Valid: true},
	})

	if err != nil {
		logger.WithFields(log.Fields{"shop_id": req.Id}).Errorf("an error in query %+v", err)
		return &pb.UpdateShopResponse{
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
			Message: "an error in query",
		}, nil
	}

	logger.WithFields(log.Fields{"shop_id": req.Id}).Info("Successfully updated shop")

	shop1 := &pb.Shop{
		Id:        shop.ID,
		Name:      shop.Name,
		Address:   shop.Address,
		OwnerId:   shop.OwnerID,
		CreatedAt: timestamppb.New(shop.CreatedAt),
		UpdatedAt: timestamppb.New(shop.UpdatedAt.Time),
	}

	return &pb.UpdateShopResponse{
		Shop:    shop1,
		Status:  http.StatusOK,
		Message: "Shop updated successfully",
	}, nil
}
