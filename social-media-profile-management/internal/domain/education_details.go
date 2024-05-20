package domain

type EducationDetail struct {
	ID            int
	Qualification string
	StartYear     int
	EndYear       int
	CourseDetails []CourseDetail
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
