package tickets

import (
	"context"
)

type Service interface {
	GetTotalTickets(ctx context.Context, destination string) (int, error)
	AverageDestination(ctx context.Context, destination string) (float64, error)
}

type service struct {
	r Repository
}

func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetTotalTickets(ctx context.Context, destination string) (int, error) {
	tickets, err := s.r.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	return len(tickets), nil
}

func (s *service) AverageDestination(ctx context.Context, destination string) (float64, error) {
	tickets, err := s.r.GetTicketByDestination(ctx, destination)
	if err != nil {
		return 0, err
	}
	allTickets, err := s.r.GetAll(ctx)
	if err != nil {
		return 0, err
	}
	tck := float64(len(tickets))
	all := float64(len(allTickets))
	return tck / all, nil
}
