package Gobluem

func (this *GoBluemClient)ErrLast()(string,error){
	cmd := CmdModel{"err","lastont",""}
	return execCmd(&cmd)
}

func (this *GoBluemClient)ErrCount()(string,error){
	cmd := CmdModel{"err","count",""}
	return execCmd(&cmd)
}