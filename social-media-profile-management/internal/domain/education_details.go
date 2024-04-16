package domain

type EducationDetail struct {
	EducationDetailID int
	Qualification     string
	StartYear         int
	EndYear           int
	CourseDetails     []CourseDetail
}

func (ed EducationDetail) ID() int {
	return ed.EducationDetailID
}

type CourseDetail struct {
	CourseName CourseName
	StartYear  int
	EndYear    int
}

type CourseName string

const (
	School        CourseName = "School"
	CollegeDegree CourseName = "College Degree"
	PostGraduate  CourseName = "Post Graduate"
)
