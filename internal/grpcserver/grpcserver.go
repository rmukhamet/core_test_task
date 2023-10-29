package grpcserver

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/rmukhamet/core_test_task/api/storage"
	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCService struct {
	pb.UnimplementedStorageServer
	port   string
	repo   RepositoryI
	server *grpc.Server
}

type RepositoryI interface {
	GetRetailerList(ctx context.Context) ([]model.Retailer, error)
	GetRetailerByID(ctx context.Context, retailerID string) (model.Retailer, error)
	DeleteRetailer(ctx context.Context, ID string) error
	DeleteRetailerVersion(ctx context.Context, ID string, version int) error
	History(ctx context.Context, retailerID string) ([]model.Retailer, error)
	GetRetailerVersion(ctx context.Context, retailerID string, version int) (model.Retailer, error)
}

func New(cfg *config.StorageConfig, repo RepositoryI) *GRPCService {
	server := grpc.NewServer()
	service := &GRPCService{
		server: server,
		port:   cfg.GRPC.Port,
		repo:   repo,
	}
	pb.RegisterStorageServer(server, service)

	return service
}

func (gs *GRPCService) Run(ctx context.Context) error {
	listen, err := net.Listen("tcp", fmt.Sprintf(":%s", gs.port))
	if err != nil {
		return fmt.Errorf("failed to listen on %s grps server %w", gs.port, err)
	}

	err = gs.server.Serve(listen)
	if err != nil {
		return fmt.Errorf("failed to serve grps server %w", err)
	}

	log.Print("gRPC server run at ", gs.port)
	return nil
}

func (gs *GRPCService) Close(_ context.Context) error {
	gs.server.GracefulStop()

	return nil
}
func (gs *GRPCService) GetRetailerList(ctx context.Context, _ *pb.Empty) (*pb.Retailers, error) {
	retailers, err := gs.repo.GetRetailerList(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed get from db: %w", err)
	}

	result := make([]*pb.Retailer, 0, len(retailers))

	for _, r := range retailers {
		result = append(result, NewRetailer(r))
	}

	return &pb.Retailers{Items: result}, nil
}
func (gs *GRPCService) GetRetailerByID(ctx context.Context, in *pb.RetailerID) (*pb.Retailer, error) {
	retailer, err := gs.repo.GetRetailerByID(ctx, in.GetID())
	if err != nil {
		return nil, fmt.Errorf("failed get from db: %w", err)
	}

	return NewRetailer(retailer), nil
}

func (gs *GRPCService) DeleteRetailer(ctx context.Context, in *pb.RetailerID) (*pb.Empty, error) {
	err := gs.repo.DeleteRetailer(ctx, in.GetID())
	if err != nil {
		return nil, fmt.Errorf("failed delete from db: %w", err)
	}

	return nil, nil
}

func (gs *GRPCService) DeleteVersion(ctx context.Context, in *pb.RetailerIDVersionID) (*pb.Empty, error) {
	err := gs.repo.DeleteRetailerVersion(ctx, in.GetRetailerID(), int(in.GetVersionID()))
	if err != nil {
		return nil, fmt.Errorf("failed delete from db: %w", err)
	}

	return nil, nil
}
func (gs *GRPCService) History(ctx context.Context, in *pb.RetailerID) (*pb.Retailers, error) {
	retailers, err := gs.repo.History(ctx, in.GetID())
	if err != nil {
		return nil, fmt.Errorf("failed get retailer version history from db: %w", err)
	}

	result := make([]*pb.Retailer, 0, len(retailers))

	for _, r := range retailers {
		result = append(result, NewRetailer(r))
	}

	return &pb.Retailers{Items: result}, nil
}
func (gs *GRPCService) GetRetailerVersion(ctx context.Context, in *pb.RetailerIDVersionID) (*pb.Retailer, error) {
	retailer, err := gs.repo.GetRetailerVersion(ctx, in.RetailerID, int(in.GetVersionID()))
	if err != nil {
		return nil, fmt.Errorf("failed get from db: %w", err)
	}

	return NewRetailer(retailer), nil
}

func NewRetailer(r model.Retailer) *pb.Retailer {
	return &pb.Retailer{
		ID:             r.ID,
		Name:           r.Name,
		AddressCity:    r.Address.City,
		AddressStreet:  r.Address.City,
		AddressHouse:   r.Address.City,
		OwnerFirstName: r.Address.City,
		OwnerLastName:  r.Address.City,
		OpenTime:       timestamppb.New(r.OpenTime),
		CloseTime:      timestamppb.New(r.CloseTime),
		Version:        int64(r.Version.Version),
		Actor:          r.Version.Actor,
		CreatedAt:      timestamppb.New(r.Version.CreatedAt),
		UpdatedAt:      timestamppb.New(r.Version.UpdatedAt),
	}
}
