package domain

type Privacy struct {
	ProfileID int
	Level     PrivacyLevel
}

func (p Privacy) ID() int {
	return p.ProfileID
}

type PrivacyLevel string

const (
	FullyPublic     PrivacyLevel = "Fully Public"
	PartiallyPublic PrivacyLevel = "Partially Public"
	Private         PrivacyLevel = "Private"
)
