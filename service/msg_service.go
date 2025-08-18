package service

import "go-micro/model"

type MsgService interface {
	Insert(*model.Message) error
	Get(*model.Message) (*model.Message, error)
	Delete(*model.Message) error
}

type msgService struct {
	repo model.MessageRepo
}

func NewMsgService(repo model.MessageRepo) MsgService {
	return &msgService{repo}
}

func (s *msgService) Insert(model *model.Message) error {
	return s.repo.Insert(model)
}
func (s *msgService) Get(model *model.Message) (*model.Message, error) {
	err := s.repo.Get(model)
	if err != nil {
		return nil, err
	}
	return model, nil
}
func (s *msgService) Delete(model *model.Message) error {
	return s.repo.Delete(model)
}
