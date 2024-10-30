package services

import (
	pb "github.com/asadlive84/bizspace/proto/shop/pb"
	q "github.com/asadlive84/bizspace/shop-svc/internal/query"
)

type Server struct {
	// H db.DbHandler
	Q q.Query
	pb.UnimplementedShopServiceServer
}
