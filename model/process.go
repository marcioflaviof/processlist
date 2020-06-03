package model

//Processes model for adding in database
type Processes struct {
	Process []string `json:"process"`
}

//AddProcess append process in slice of processes
func (p *Processes) AddProcess(process string) {
	p.Process = append(p.Process, process)
}
