package domain

type WorkDetail struct {
	WorkDetailID   int
	EmployerName   string
	StartYear      int
	EndYear        int
	JobRole        string
	JobDescription string
}

func (wd WorkDetail) ID() int {
	return wd.WorkDetailID
}
