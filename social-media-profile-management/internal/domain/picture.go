package domain

// Picture represents a picture of a user
type Picture struct {
	ID          int         `json:"pictureID"`
	PictureURL  string      `json:"pictureURL"`
	Size        int         `json:"size"`
	PictureType PictureType `json:"pictureType"`
}

// PictureType represents the type of picture
type PictureType string

// Picture types
const (
	// CoverPhoto represents a cover photo
	CoverPhoto PictureType = "cover"
	// ProfilePhoto represents a profile photo
	ProfilePhoto PictureType = "profile"
)
