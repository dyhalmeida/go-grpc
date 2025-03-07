package service

import (
	"context"

	"github.com/dyhalmeida/go-grpc/internal/database"
	"github.com/dyhalmeida/go-grpc/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, request *pb.CategoryRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Create(request.Name, request.Description)
	if err != nil {
		return nil, err
	}

	return &pb.Category{Id: category.ID, Name: category.Name, Description: category.Description}, nil
}
