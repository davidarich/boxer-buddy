package domain

// System process
type Process struct {
	Id       int
	ParentId int
	Path     string
	FileName string
}

/*
Go uses zero values by default which can refer to a valid process (system/kernel process)
Process with ID set to -1 indicates the struct has not been initialized with data
*/
func NewProcess() *Process {
	return &Process{
		Id:       -1,
		ParentId: -1,
	}
}
