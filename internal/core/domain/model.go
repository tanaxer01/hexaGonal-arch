package domain

type FeatureFlagResponse interface {
	IsFeatureFlagResponse()
}

type FeatureFlagT struct {
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Enabled     bool   `json:"enabled" bson:"enabled"`
}

func (FeatureFlagT) IsFeatureFlagResponse() {}

type DBErrorT struct {
	Message string `json:"message"`
}

func (DBErrorT) IsFeatureFlagResponse() {}

type FeatureFlagInput struct {
	Name string `json:"name"`
}

type Query struct{}
