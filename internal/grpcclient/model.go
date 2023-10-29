package grpcclient

import (
	pb "github.com/rmukhamet/core_test_task/api/storage"
	"github.com/rmukhamet/core_test_task/internal/model"
)

func RetailerToDTO(r *pb.Retailer) model.Retailer {
	return model.Retailer{
		ID:   r.GetID(),
		Name: r.GetName(),
		Address: model.Address{
			City:   r.GetAddressCity(),
			Street: r.GetAddressStreet(),
			House:  r.GetAddressHouse(),
		},
		Owner: model.Person{
			FirstName: r.GetOwnerFirstName(),
			LastName:  r.GetOwnerLastName(),
		},
		OpenTime:  r.OpenTime.AsTime(),
		CloseTime: r.CloseTime.AsTime(),
		Version: model.Version{
			Version:   int(r.GetVersion()),
			Actor:     r.GetActor(),
			CreatedAt: r.CreatedAt.AsTime(),
			UpdatedAt: r.UpdatedAt.AsTime(),
		},
	}
}
