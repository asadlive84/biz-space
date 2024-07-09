package services

import (
	"context"
	"os"
	"time"

	// "log"
	"net/http"

	pb "github.com/asadlive84/bizspace/proto/shop/pb"

	log "github.com/sirupsen/logrus"
)

func (s *Server) DeleteShop(ctx context.Context, req *pb.DeleteShopRequest) (*pb.DeleteShopResponse, error) {
	var logger = log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})
	logger.SetOutput(os.Stdout)

	logger.WithFields(log.Fields{"shop_id": req.Id}).Info("Service: DeleteShop method")

	err := s.Q.DeleteShop(ctx, req.Id)
	if err != nil {
		logger.WithFields(log.Fields{"shop_id": req.Id}).Errorf("an error in delete query %+v", err)
		return &pb.DeleteShopResponse{
			Status:  http.StatusBadRequest,
			Error:   err.Error(),
			Message: "an error in query",
		}, nil
	}

	logger.WithFields(log.Fields{"shop_id": req.Id}).Info("Successfully deleted shop")

	return &pb.DeleteShopResponse{
		Status:  http.StatusOK,
		Message: "Shop deleted successfully",
	}, nil
}
