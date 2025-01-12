package gapi

import (
	"fmt"

	db "github.com/PetrusZ/simplebank/db/sqlc"
	"github.com/PetrusZ/simplebank/pb"
	"github.com/PetrusZ/simplebank/token"
	"github.com/PetrusZ/simplebank/util"
	"github.com/PetrusZ/simplebank/worker"
)

// Server serves gRPC requests for our banking service.
type Server struct {
	pb.UnimplementedSimpleBankServer
	config          util.Config
	store           db.Store
	tokenMaker      token.Maker
	taskDistributor worker.TaskDistributor
}

// NewServer creates a new gRPC server.
func NewServer(config util.Config, store db.Store, taskDistributor worker.TaskDistributor) (*Server, error) {
	tokerMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey)
	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %w", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokerMaker,
		taskDistributor: taskDistributor,
	}

	return server, nil
}
