package domain

type Privacy struct {
	ID    int
	Level PrivacyLevel
}

type PrivacyLevel string

const (
	FullyPublic     PrivacyLevel = "Fully Public"
	PartiallyPublic PrivacyLevel = "Partially Public"
	Private         PrivacyLevel = "Private"
)
