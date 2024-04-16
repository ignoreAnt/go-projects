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

func (up UserProfile) ID() int {
	return up.ProfileID
}
