package domain

type UserProfile struct {
	ProfileID      int
	FirstName      string
	MiddleName     string
	LastName       string
	Location       string
	Sex            string
	Age            int
	WorkDetails    []WorkDetail
	Education      []EducationDetail
	RelationStatus RelationStatus
	ProfileType    ProfileType
}

type RelationStatus string

const (
	Single      RelationStatus = "single"
	Married     RelationStatus = "married"
	Engaged     RelationStatus = "engaged"
	Complicated RelationStatus = "complicated"
)

type ProfileType string

const (
	BusinessProfile ProfileType = "business"
	Public          ProfileType = "public"
)

func (up UserProfile) ID() int {
	return up.ProfileID
}
