package tm1

import (
	"tm1-api/domain"
)

type Service interface {
	GetTm(uri1 string, uri2 string, queryString string) (any, error)
	SendTm(input domain.Tm1RequestData) (any, error)
	SendRaTest(input domain.Tm1RequestDynamicData) (any, error)
	AddElementTm(input domain.Tm1AddElementRequestData) (any, error)
	SendDynamicTm(input domain.Tm1DynamicRequestData) (any, error)
	SendGetDynamicTm(input domain.Tm1GetElementRequestData) (any, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) AddElementTm(input domain.Tm1AddElementRequestData) (any, error) {
	res, err := s.repository.AddElementTm(input)

	return res, err
}

func (s *service) SendDynamicTm(input domain.Tm1DynamicRequestData) (any, error) {
	res, err := s.repository.SendDynamicTm(input)

	return res, err
}

func (s *service) SendGetDynamicTm(input domain.Tm1GetElementRequestData) (any, error) {
	res, err := s.repository.SendGetDynamicTm(input)

	return res, err
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
