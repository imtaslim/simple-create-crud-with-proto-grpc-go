package categories

import (
	"context"
	"log"
	"time"

	cpb "grpc/proto/category"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct{}

type Categories struct{
	ID int64
	Name string
	Description string
	Status bool
}
var categories = []Categories{
	{
		ID: 1,
		Name: "Mystry",
		Description: "Something that is difficult or impossible to understand or explain.",
		Status: true,
	},
	{
		ID: 2,
		Name: "Horror",
		Description: "An intense feeling of fear, shock, or disgust.",
		Status: true,
	},
	{
		ID: 3,
		Name: "Romance",
		Description: "A quality or feeling of mystery, excitement, and remoteness associate with love.",
		Status: false,
	},
}

func (s *Server) GetCategory(ctx context.Context, req *cpb.GetCategoryRequest) (*cpb.GetCategoryResponse, error) {
	log.Printf("Category ID: %d", req.GetID())
	var cat Categories
	for _, c := range categories {
		if c.ID == req.GetID() {
			cat = c
			break
		}
	}

	if cat.ID == 0 {
		return &cpb.GetCategoryResponse{}, status.Errorf(codes.NotFound, "invalid id")
	}

	return &cpb.GetCategoryResponse{
		ID: cat.ID,
		Name: cat.Name,
		Description: cat.Description,
		Status: cat.Status,
	}, nil
}

func (s *Server) GetCategories(req *cpb.GetCategoriesRequest, stream cpb.CategoryService_GetCategoriesServer) error{
	for _, c := range categories {
		err := stream.Send(&cpb.GetCategoryResponse{
			ID: c.ID,
			Name: c.Name,
			Description: c.Description,
			Status: c.Status,
		})
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to send Categories")
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}

func (s *Server) GetCat(stream cpb.CategoryService_GetCatServer) error{
	for _, c := range categories {
		err := stream.SendAndClose(&cpb.GetCategoryResponse{
			ID: c.ID,
			Name: c.Name,
			Description: c.Description,
			Status: c.Status,
		})
		if err != nil {
			return status.Errorf(codes.Internal, "Failed to send Categories")
		}
		time.Sleep(time.Second * 2)
	}
	return nil
}
