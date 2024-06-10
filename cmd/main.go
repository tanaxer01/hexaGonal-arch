package main

import (
	"fmt"
	"hexago/internal/adapters/repository"
	"hexago/internal/core/services"
)

var svc *services.FeatureFlagService

func main() {
	store := repository.NewMongoDBRepository("ms-feature-flags-api", "feature-flags")
	svc = services.NewFeatureFlagService(store)

	res, err := svc.GetFeatureFlag("This")
	if err != nil {
	}

	fmt.Printf("--> %+v\n", res)
}
