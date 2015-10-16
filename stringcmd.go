package Gobluem
import (
)

func (this *GoBluemClient)Add(key ,value string)(string,error){
	cmd := CmdModel{Cmd:"add",Key:key,Value:value}
	return execCmd(&cmd)
}

func (this *GoBluemClient)Get(key  string)(string,error){
	cmd := CmdModel{Cmd:"get",Key:key}
	return execCmd(&cmd)
}


func (this *GoBluemClient)Set(key ,value string)(string,error){
	cmd := CmdModel{Cmd:"set",Key:key,Value:value}
	return execCmd(&cmd)
}

func (this *GoBluemClient)Del(key  string)(string,error){
	cmd := CmdModel{Cmd:"del",Key:key,Value:""}
	return execCmd(&cmd)
}

func (this *GoBluemClient)Append(key ,value string)(string,error){
	cmd := CmdModel{Cmd:"append",Key:key,Value:value}
	return execCmd(&cmd)
}


func (this *GoBluemClient)GetSet(key ,value string)(string,error){
	cmd := CmdModel{Cmd:"getset",Key:key,Value:value}
	return execCmd(&cmd)
}