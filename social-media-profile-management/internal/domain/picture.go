package domain

// Picture represents a picture of a user
type Picture struct {
	PictureID   int
	PictureURL  string
	Size        int
	PictureType PictureType
}

// ID returns the ID of the picture
func (p Picture) ID() int {
	return p.PictureID
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
