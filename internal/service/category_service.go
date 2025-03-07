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

func (c *CategoryService) ListCategories(ctx context.Context, request *pb.Null) (*pb.CategoryList, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}

	var categoriesPb []*pb.Category

	for _, category := range categories {
		categoriesPb = append(categoriesPb, &pb.Category{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}

	return &pb.CategoryList{Categories: categoriesPb}, nil
}

func (c *CategoryService) GetCategoryById(ctx context.Context, request *pb.CategoryByIdRequest) (*pb.Category, error) {
	category, err := c.CategoryDB.Find(request.Id)
	if err != nil {
		return nil, err
	}

	return &pb.Category{Id: category.ID, Name: category.Name, Description: category.Description}, nil
}
