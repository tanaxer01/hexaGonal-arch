package handler

import (
	"hexago/internal/core/domain"
	"hexago/internal/core/services"
)

type FeatureFlagHandler struct {
	svc services.FeatureFlagService
}

func NewFeatureFlagHandler(FeatureFlagService services.FeatureFlagService) *FeatureFlagHandler {
	return &FeatureFlagHandler{
		svc: FeatureFlagService,
	}
}

func (h *FeatureFlagHandler) GetFeatureFlag(name string) domain.FeatureFlagResponse {
	featureFlag, err := h.svc.GetFeatureFlag(name)
	if err != nil {
		return domain.DBErrorT{Message: err.Error()}
	}

	return featureFlag
}
