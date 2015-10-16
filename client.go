package Gobluem

import (
	"bufio"
	"net"
	"encoding/json"
)

var continuechan chan bool
var serverHost string
var connObj *net.TCPConn

type GoBluemClient struct {

}

func (this *GoBluemClient)Connection(serverHost string)error{
	err:= beginConn(serverHost)
	return err
}

func (this *GoBluemClient)Close()error{
	return connObj.Close()
}

//创建连接
func beginConn(host string) error {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp4", host)
	var err error
	connObj, err = net.DialTCP("tcp", nil, tcpAddr)
	return err
}
//发送命令
func sendCmd(cmd string)(int ,error){
	return connObj.Write([]byte(cmd +"\n"))
}

//接受结果
func receiveMsg()(string,error){
	reader := bufio.NewReader(connObj)
	msg, err := reader.ReadString('\n')
	return msg,err
}


func execCmd(cmd *CmdModel)(string,error){

	json,err := json.Marshal(*cmd)
	if err!=nil {
		return "",err
	}
	str := string(json)
	//发送命令
	_,err = sendCmd(str)
	if err!=nil {
		return "",err
	}
	//接受结果
	msg,err := receiveMsg()
	return msg,err
}
//
//func main() {
//	cmd()
//	var tcpAddr *net.TCPAddr
//	tcpAddr, _ = net.ResolveTCPAddr("tcp4", serverHost		)
//
//	conn, _ := net.DialTCP("tcp", nil, tcpAddr)
//	defer conn.Close()
//	fmt.Println("connected!")
//
//	go onMessageRecived(conn)
//	var s string
//	continuechan = make(chan bool, 1)
//
//	continuechan <- true
//
//	for {
//		<-continuechan
//		fmt.Print("bule-clent: ")
//		s = consoleReadline()
//		conn.Write([]byte(s + "\n"))
//	}
//}
//
//func onMessageRecived(conn *net.TCPConn) {
//	reader := bufio.NewReader(conn)
//	for {
//		msg, err := reader.ReadString('\n')
//		fmt.Print(msg)
//		if err != nil {
//			fmt.Println(err.Error())
//		}
//		continuechan <- true
//	}
//}
//
//func consoleReadline()string{
//	reader := bufio.NewReader(os.Stdin)
//	b,_,_ := reader.ReadLine()
//	return string(b)
//}
//
//
//func cmd(){
//	port := flag.Int("port", 8090, "bluemdb server port")
//	ip := flag.String("ip","127.0.0.1","bluemdb server ip address")
//
//	flag.Parse()
//
//	serverHost = fmt.Sprintf("%s:%d",*ip,*port)
//
//}