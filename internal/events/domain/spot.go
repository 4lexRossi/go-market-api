package domain

import (
	"errors"

	"github.com/google/uuid"
)

var (
	ErrInvalidSpotNumber       = errors.New("invalid spot number")
	ErrSpotNotFound            = errors.New("spot not found")
	ErrSpotAlreadyReserved     = errors.New("spot already reserved")
	ErrSpotMinLenght           = errors.New("spot must be at least 2 characters long")
	ErrSpotMustStartWithLetter = errors.New("spot must start with a character")
	ErrSpotMustEndWithNumber   = errors.New("spot must end with a number")
)

type SpotStatus string

const (
	SpotStatusAvailable SpotStatus = "available"
	SpotStatusSold      SpotStatus = "sold"
)

type Spot struct {
	ID       string
	EventID  string
	Name     string
	Status   SpotStatus
	TicketID string
}

func NewSpot(event *Event, name string) (*Spot, error) {
	spot := &Spot{
		ID:      uuid.New().String(),
		EventID: event.ID,
		Name:    name,
		Status:  SpotStatusAvailable,
	}

	if err := spot.ValidateSpot(); err != nil {
		return nil, err
	}
	return spot, nil
}

func (s Spot) ValidateSpot() error {
	if s.Name == "" {
		return ErrInvalidSpotNumber
	}

	if len(s.Name) < 2 {
		return ErrSpotMinLenght
	}

	if s.Name[0] < 'A' || s.Name[0] > 'Z' {
		return ErrSpotMustStartWithLetter
	}

	if s.Name[1] < '0' || s.Name[1] > '9' {
		return ErrSpotMustEndWithNumber
	}
	return nil
}

func (s *Spot) Reserve(ticketID string) error {
	if s.Status == SpotStatusSold {
		return ErrSpotAlreadyReserved
	}

	s.Status = SpotStatusSold
	s.TicketID = ticketID
	return nil
}
