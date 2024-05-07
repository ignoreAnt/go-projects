package memory

import (
	"social-server/internal/domain"
	"testing"
)

func setup(t *testing.T) domain.Picture {
	t.Helper()
	pictureType := domain.CoverPhoto
	picture := domain.Picture{
		PictureURL:  "https://s3.aws.com/pics/123456",
		Size:        100,
		PictureType: pictureType,
	}

	return picture

}

func Test_NewPictureRepository(t *testing.T) {
	// Create memoenry map of pictures
	pics := map[int]domain.Picture{}

	want := PictureRepository{}
	want.pictures = pics
	want.nextID = 1

	got := NewPictureRepository()

	if got.nextID != want.nextID {
		t.Errorf("got %v, want %v", got.nextID, want.nextID)
	}

	if len(got.pictures) != len(want.pictures) {
		t.Errorf("got %v, want %v", len(got.pictures), len(want.pictures))
	}

}

func Test_Create(t *testing.T) {
	picture := setup(t)
	picture.PictureID = 1

	want := PictureRepository{
		pictures: map[int]domain.Picture{1: picture},
		nextID:   2,
	}

	got := NewPictureRepository()
	err := got.Create(picture)
	if err != nil {
		t.Error("Error Creating Picture Repo")
	}

	// t.Log("got", got.pictures, "want", want.pictures)

	if want.nextID != got.nextID {
		t.Errorf("got %v, want %v", got.nextID, want.nextID)
	}

	if len(want.pictures) != len(got.pictures) {
		t.Errorf("got %v, want %v", len(got.pictures), len(want.pictures))
	}

	if ok := got.pictures[1]; ok != picture {
		t.Errorf("got %v, want %v", ok, picture)
	}

}

func Test_Update(t *testing.T) {
	picture1 := setup(t)
	picture1.PictureID = 1

	updatedPic := setup(t)
	updatedPic.PictureID = 1
	updatedPic.Size = 200

	// want := PictureRepository{
	// 	pictures: map[int]domain.Picture{1: updatedPic},
	// 	nextID:   2,
	// }

	got := NewPictureRepository()
	got.Create(picture1)
	got.Update(updatedPic)

	if ok := got.pictures[1]; ok != updatedPic {
		t.Errorf("got %v, want %v", ok, updatedPic)
	}

	if ok := got.pictures[1].Size == 200; !ok {
		t.Errorf("got %v, want %v", ok, updatedPic.Size)
	}

}

func Test_Delete(t *testing.T) {
	picture := setup(t)
	picture.PictureID = 1

	got := NewPictureRepository()
	err := got.Create(picture)
	if err != nil {
		t.Error("Error Creating Picture Repo")
	}

	got.Delete(1)

	if len(got.pictures) != 0 {
		t.Errorf("got %v, want %v", len(got.pictures), 0)
	}

}

func Test_GetById(t *testing.T) {

	picture := setup(t)
	picture.PictureID = 1

	got := NewPictureRepository()
	err := got.Create(picture)
	if err != nil {
		t.Error("Error Creating Picture Repo")
	}

	pic, err := got.GetById(1)

	if err != nil {
		t.Errorf("Failed to get Picture by Id Error: %v", err)
	}

	if pic.PictureID != 1 {
		t.Errorf("got %v, want %v", pic.PictureID, 1)
	}

}

func Test_CreateV2(t *testing.T) {
	picture := setup(t)
	picture.PictureID = 1

	repo := NewPictureRepository()

	t.Run("Testing Creating Picture", func(t *testing.T) {
		err := repo.Create(picture)
		if err != nil {
			t.Errorf("Error Creating picture: %v", err)
		}

		if len(repo.pictures) != 1 {
			t.Errorf("Picture creation failed got: %v want: %v", len(repo.pictures), 1)
		}
	})

}
