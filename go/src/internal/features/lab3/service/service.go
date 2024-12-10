package service

import "context"

type Service struct{}

func (s *Service) Execute(ctx context.Context, in *In) (*Out, error) {
	// TODO implement for lab 3
	return &Out{}, nil
}

func New() *Service {
	return &Service{}
}
