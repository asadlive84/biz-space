package services

import (
	q "github.com/asadlive84/bizspace/auth-svc/internal/query"
	pb "github.com/asadlive84/bizspace/proto/auth/pb"
)

type Server struct {
	// H db.DbHandler
	Q q.Query
	pb.UnimplementedAuthServiceServer
}
