package grpcclient

import (
	"context"
	"fmt"
	"log"

	pb "github.com/rmukhamet/core_test_task/api/storage"
	"github.com/rmukhamet/core_test_task/internal/config"
	"github.com/rmukhamet/core_test_task/internal/model"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GRPCClient struct {
	conn   *grpc.ClientConn
	client pb.StorageClient
}

func New(cfg *config.GRPC) *GRPCClient {
	conn, err := grpc.Dial(
		fmt.Sprintf("%s:%s", cfg.Addr, cfg.Port),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	c := pb.NewStorageClient(conn)

	return &GRPCClient{
		conn:   conn,
		client: c,
	}
}

func (c *GRPCClient) Close(_ context.Context) error {
	return c.conn.Close()
}
func (c *GRPCClient) GetRetailerList(ctx context.Context) ([]model.Retailer, error) {
	retailers, err := c.client.GetRetailerList(ctx, &pb.Empty{})
	if err != nil {
		return nil, err
	}

	result := make([]model.Retailer, 0, len(retailers.Items))

	for _, r := range retailers.Items {
		retailer := RetailerToDTO(r)
		result = append(result, retailer)
	}

	return result, nil
}
func (c *GRPCClient) GetRetailerByID(ctx context.Context, retailerID string) (model.Retailer, error) {
	retailer, err := c.client.GetRetailerByID(ctx, &pb.RetailerID{ID: retailerID})
	if err != nil {
		return model.Retailer{}, err
	}

	return RetailerToDTO(retailer), nil
}

func (c *GRPCClient) DeleteRetailer(ctx context.Context, retailerID string) error {
	_, err := c.client.DeleteRetailer(ctx, &pb.RetailerID{ID: retailerID})

	return err
}

func (c *GRPCClient) GetRetailerVersion(ctx context.Context, retailerID string, versionID int) (model.Retailer, error) {
	retailer, err := c.client.GetRetailerVersion(ctx, &pb.RetailerIDVersionID{
		RetailerID: retailerID,
		VersionID:  int64(versionID),
	})
	if err != nil {
		return model.Retailer{}, err
	}

	return RetailerToDTO(retailer), nil

}

func (c *GRPCClient) DeleteRetailerVersion(ctx context.Context, retailerID string, versionID int) error {
	_, err := c.client.DeleteRetailerVersion(ctx, &pb.RetailerIDVersionID{
		RetailerID: retailerID,
		VersionID:  int64(versionID),
	})

	return err
}

func (c *GRPCClient) GetRetailerVersionList(ctx context.Context, retailerID string) ([]model.Retailer, error) {
	retailers, err := c.client.History(ctx, &pb.RetailerID{ID: retailerID})
	if err != nil {
		return nil, err
	}

	result := make([]model.Retailer, 0, len(retailers.Items))

	for _, r := range retailers.Items {
		retailer := RetailerToDTO(r)
		result = append(result, retailer)
	}

	return result, nil
}
