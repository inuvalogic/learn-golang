package services

import (
	"kasir-api/models"
	"kasir-api/repositories"
)

type ProductService struct {
	repo *repositories.ProductRepository
}

func NewProductService(repo *repositories.ProductRepository) *ProductService {
	return &ProductService{repo: repo}
}

func (s *ProductService) GetAll(name string) ([]models.Product, error) {
	return s.repo.GetAll(name)
}

func (s *ProductService) Create(req models.CreateProductRequest) (*models.Product, error) {
	product := &models.Product{
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
		Category: models.Category{
			ID: req.CategoryID,
		},
	}

	if err := s.repo.Create(product); err != nil {
		return nil, err
	}

	if err := s.repo.LoadCategory(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) GetByID(id int) (*models.Product, error) {
	return s.repo.GetByID(id)
}

func (s *ProductService) Update(id int, req models.UpdateProductRequest) (*models.Product, error) {
	product := &models.Product{
		ID:    id,
		Name:  req.Name,
		Price: req.Price,
		Stock: req.Stock,
		Category: models.Category{
			ID: req.CategoryID,
		},
	}

	if err := s.repo.Update(product); err != nil {
		return nil, err
	}

	if err := s.repo.LoadCategory(product); err != nil {
		return nil, err
	}

	return product, nil
}

func (s *ProductService) Delete(id int) error {
	return s.repo.Delete(id)
}