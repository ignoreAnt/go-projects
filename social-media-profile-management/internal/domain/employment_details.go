package domain

type EmploymentDetail struct {
	EmploymentDetailID int
	EmployerName       string
	StartYear          int
	EndYear            int
	JobRole            string
	JobDescription     string
}

func (ed EmploymentDetail) ID() int {
	return ed.EmploymentDetailID
}
