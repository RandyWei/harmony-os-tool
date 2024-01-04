package models

type Process struct {
	PID  string `json:"pid"`
	VIRT string `json:"virt"`
	RES  string `json:"res"`
	SHR  string `json:"shr"`
	CPU  string `json:"cpu"`
	MEM  string `json:"mem"`
	TIME string `json:"time"`
	ARGS string `json:"args"`
	DESC string `json:"desc"`
}

type TopInfo struct {
	Processes []Process `json:"processes"`
	Tasks     string    `json:"tasks"`
	Mem       string    `json:"mem"`
}
