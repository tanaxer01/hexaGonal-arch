package ports

import "hexago/internal/core/domain"

type FeatureFlagService interface {
	GetFeatureFlag(name string) (*domain.FeatureFlagT, error)
}

type FeatureFlagRepository interface {
	GetFeatureFlag(name string) (*domain.FeatureFlagT, error)
}
