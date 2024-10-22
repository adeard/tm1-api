package tm1

import "tm1-api/domain"

type Service interface {
	GetTm(uri1 string, uri2 string, queryString string) (any, error)
	SendTm(input domain.Tm1RequestData) (any, error)
	SendRaTest(input domain.Tm1RequestDynamicData) (any, error)
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

func (s *service) SendRaTest(input domain.Tm1RequestDynamicData) (any, error) {
	res, err := s.repository.SendRaTest(input)

	return res, err
}

func (s *service) GetTm(uri1 string, uri2 string, queryString string) (any, error) {
	res, err := s.repository.GetTm(uri1, uri2, queryString)

	return res, err
}
