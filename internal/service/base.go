package service

import (
	"backend-vtb/internal/repository"
	"log/slog"

	"github.com/google/uuid"
)

type BaseService struct {
	repos  *repository.Repository
	logger *slog.Logger
}

func NewBaseService(repos *repository.Repository, logger *slog.Logger) *BaseService {
	return &BaseService{
		repos:  repos,
		logger: logger,
	}
}
func (s *BaseService) GetAchievements(id uuid.UUID) ([]string, error) {
	return nil, nil
}

func (s *BaseService) GetName(id uuid.UUID) (string, error) {
	return "", nil
}

func (s *BaseService) GetAmount(id uuid.UUID) (int, error) {
	return 0, nil
}

func (s *BaseService) GetBaseInfo(id uuid.UUID) (string, error) {
	return "", nil
}

func (s *BaseService) GetNeuroMean(id uuid.UUID) (float64, error) {
	return 0, nil
}

func (s *BaseService) GetCryptoData(id uuid.UUID) (string, error) {
	return "", nil
}

func (s *BaseService) GetAPIInfo(id uuid.UUID) (string, error) {
	return "", nil
}

func (s *BaseService) GetFullAPIInfo(id uuid.UUID) (string, error) {
	return "", nil
}

func (s *BaseService) GetFines(id uuid.UUID) ([]int, error) {
	return nil, nil
}

func (s *BaseService) GetFineByID(id uuid.UUID) (int, error) {
	return 0, nil
}

func (s *BaseService) GetPayments(id uuid.UUID) ([]int, error) {
	return nil, nil
}

func (s *BaseService) GetPaymentByID(id uuid.UUID) (int, error) {
	return 0, nil
}

func (s *BaseService) GetStatsData(id uuid.UUID) (string, error) {
	return "", nil
}

func (s *BaseService) GetAnalyze(id uuid.UUID) (string, error) {
	return "", nil
}
