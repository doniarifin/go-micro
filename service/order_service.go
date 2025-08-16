package service

import (
	"go-micro/model"
)

type OrderService interface {
	Create(model *model.Order) error
	Get(model *model.Order) (*model.Order, error)
	Update(model *model.Order) error
	Delete(model *model.Order) error
}

type orderService struct {
	repo model.OrderRepository
}

func NewOrderService(repo model.OrderRepository) OrderService {
	return &orderService{repo}
}

func (s *orderService) Create(model *model.Order) error {
	return s.repo.Insert(model)
}

func (s *orderService) Get(model *model.Order) (*model.Order, error) {
	order, err := s.repo.Read(model)
	return order, err
}

func (s *orderService) Update(model *model.Order) error {
	return s.repo.Insert(model)
}

func (s *orderService) Delete(model *model.Order) error {
	return s.repo.Delete(model)
}
