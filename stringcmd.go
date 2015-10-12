package Gobluem
import "fmt"

func (this *GoBluemClient)Add(key ,value string)(string,error){
	cmd := fmt.Sprintf("add %s %s",key,value)
	return execCmd(cmd)
}

func (this *GoBluemClient)Get(key  string)(string,error){
	cmd := fmt.Sprintf("get %s",key)
	return execCmd(cmd)
}


func (this *GoBluemClient)Set(key ,value string)(string,error){
	cmd := fmt.Sprintf("set %s %s",key,value)
	return execCmd(cmd)
}

func (this *GoBluemClient)Del(key  string)(string,error){
	cmd := fmt.Sprintf("del %s",key)
	return execCmd(cmd)
}

func (this *GoBluemClient)Append(key ,value string)(string,error){
	cmd := fmt.Sprintf("append %s %s",key,value)
	return execCmd(cmd)
}


func (this *GoBluemClient)GetSet(key ,value string)(string,error){
	cmd := fmt.Sprintf("getset %s %s",key,value)
	return execCmd(cmd)
}