package Gobluem

type CmdModel struct {
	Cmd   string
	Key   string
	Value string
}

type ReturnModel struct  {
	Status int
	Message string
	Data interface{}
}