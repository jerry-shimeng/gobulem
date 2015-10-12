package Gobluem

func (this *GoBluemClient)ErrLast()(string,error){
	cmd := ("err last")
	return execCmd(cmd)
}

func (this *GoBluemClient)ErrCount()(string,error){
	cmd := ("err count")
	return execCmd(cmd)
}