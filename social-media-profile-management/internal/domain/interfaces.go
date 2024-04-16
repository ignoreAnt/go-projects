package domain

// UserRepository User Management
type UserRepository interface {
	Create(user User) (int, error)
	Update(user User) error
	GetById(id int) (*User, error)
	Delete(id int) error
	UpdateUserName(userID int, newUserName string) error
}

// UserProfileRepository UserProfile Management
type UserProfileRepository interface {
	Create(profile UserProfile) error
	Update(profile UserProfile) error
	GetById(id int) (*UserProfile, error)
	Delete(id int) error
}

// PictureRepository Picture Management
type PictureRepository interface {
	Create(picture Picture) error
	Update(picture Picture) error
	Delete(pictureID int) error
	GetById(pictureID int) (*Picture, error)
}

// PrivacyRepository Privacy Settings Management
type PrivacyRepository interface {
	SetPrivacy(profileID int, privacy Privacy) error
	CreatePrivacy(privacy Privacy) error
}

// WorkDetailsRepository WorkDetails Management
type WorkDetailsRepository interface {
	Create(workDetail WorkDetail) error
	Update(workDetail WorkDetail) error
	Delete(workDetailsID int) error
}

// EmploymentDetailsRepository Employment Details Management
type EmploymentDetailsRepository interface {
	Create(employmentDetail EmploymentDetail) error
	Update(employmentDetail EmploymentDetail) error
	Delete(employmentDetailsID int) error
	Get(employmentDetailsID int) (*EmploymentDetail, error)
}

// EducationDetailsRepository Education Details Management
type EducationDetailsRepository interface {
	Create(educationDetails EducationDetail) error
	Update(educationDetails EducationDetail) error
	Delete(educationDetailsID int) error
}

// LifeEventsRepository Life Events Management
type LifeEventsRepository interface {
	Create(lifeEvent LifeEvent) error
	Update(lifeEvent LifeEvent) error
	Delete(lifeEventID int) error
}

// DataManager Generic DataManager for shared functionality
type DataManager interface {
	Create(data interface{}) error
	Update(data interface{}) error
	Delete(dataID int) error
}
