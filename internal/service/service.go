package service

import (
	"backend-vtb/internal/repository"
	"log/slog"

	"github.com/google/uuid"
)

type Base interface {
	GetName(id uuid.UUID) (string, error)
	GetAmount(id uuid.UUID) (int, error)
	GetAchievements(id uuid.UUID) ([]string, error)
	GetBaseInfo(id uuid.UUID) (string, error)
	GetNeuroMean(id uuid.UUID) (float64, error)
	GetCryptoData(id uuid.UUID) (string, error)
	GetAPIInfo(id uuid.UUID) (string, error)
	GetFullAPIInfo(id uuid.UUID) (string, error)
	GetFines(id uuid.UUID) ([]int, error)
	GetFineByID(id uuid.UUID) (int, error)
	GetPayments(id uuid.UUID) ([]int, error)
	GetPaymentByID(id uuid.UUID) (int, error)
	GetStatsData(id uuid.UUID) (string, error)
	GetAnalyze(id uuid.UUID) (string, error)
}

type Service struct {
	Base Base
}

func NewService(repos *repository.Repository, logger *slog.Logger) *Service {
	return &Service{
		Base: NewBaseService(repos, logger),
	}
}
