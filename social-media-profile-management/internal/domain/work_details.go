package domain

type WorkDetail struct {
	WorkDetailID   int    `json:"workDetailID"`
	EmployerName   string `json:"employerName"`
	StartYear      int    `json:"startYear"`
	EndYear        int    `json:"endYear"`
	JobRole        string `json:"jobRole"`
	JobDescription string `json:"jobDescription"`
}

func (wd WorkDetail) ID() int {
	return wd.WorkDetailID
}
