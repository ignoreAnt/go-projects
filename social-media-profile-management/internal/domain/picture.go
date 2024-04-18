package domain

type Picture struct {
	PictureID   int         `json:"pictureID"`
	PictureURL  string      `json:"pictureURL"`
	Size        int         `json:"size"`
	PictureType PictureType `json:"pictureType"`
}

func (p Picture) ID() int {
	return p.PictureID
}

type PictureType string

const (
	CoverPhoto   PictureType = "cover"
	ProfilePhoto PictureType = "profile"
)
