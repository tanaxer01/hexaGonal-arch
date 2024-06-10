package services

import (
	"hexago/internal/core/domain"
	"hexago/internal/core/ports"
	"log"
)

type FeatureFlagService struct {
	repo ports.FeatureFlagRepository
}

func NewFeatureFlagService(repo ports.FeatureFlagRepository) *FeatureFlagService {
	return &FeatureFlagService{repo: repo}
}

func (s *FeatureFlagService) GetFeatureFlag(name string) (domain.FeatureFlagResponse, error) {
	ff, err := s.repo.GetFeatureFlag(name)
	if err != nil {
		log.Printf("Error in GetFeatureFlag('%s'): %s", name, err.Error())
		return nil, err
	}

	log.Printf("Response to GetFeatureFlag('%s'): %+v", name, ff)
	return ff, nil
}
