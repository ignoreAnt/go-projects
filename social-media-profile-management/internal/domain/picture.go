package domain

type Picture struct {
	PictureID   int
	PictureURL  string
	Size        int
	PictureType PictureType
}

func (p Picture) ID() int {
	return p.PictureID
}

type PictureType string

const (
	CoverPhoto   PictureType = "cover"
	ProfilePhoto PictureType = "profile"
)
