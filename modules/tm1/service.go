package tm1

import "tm1-api/domain"

type Service interface {
	SendTm(input domain.Tm1RequestData) (any, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) SendTm(input domain.Tm1RequestData) (any, error) {
	res, err := s.repository.SendTm(input)

	return res, err
}
