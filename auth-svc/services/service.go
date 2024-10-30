package services

import (
	q "github.com/asadlive84/bizspace/auth-svc/internal/query"
	pb "github.com/asadlive84/bizspace/proto/auth/pb"
)

type Server struct {
	// H db.DbHandler
	pb.UnimplementedAuthServiceServer
	Q         q.Query
	jwtSecret string
}

func NewServer(q q.Query, jwtSecret string) *Server {
	return &Server{
		Q:         q,
		jwtSecret: jwtSecret,
	}
}
