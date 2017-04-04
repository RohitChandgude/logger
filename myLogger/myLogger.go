package myLogger
import "fmt"
import "os"
import "time"

var fp *os.File

var Level = [...]string {"INFO","WARN","ERROR","FATAL"}

const(
	Info_msg = iota
	Warn_msg
	Error_msg
	Fatal_msg
)

func Init(fileName string){
	fileInfo, err := os.Stat(fileName)
	if err != nil{
		nw, err := os.Create(fileName)
		if err != nil{
			fmt.Println("file creation err")
		}
		fmt.Println(fileInfo)
		nw.Close()
	}

	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil{
		panic(err)
	}
	fp= f
}

func logPrint(msg string, level int){
	t := time.Now()
    	for i := level; i < len(Level); i++ {
		strMsg := fmt.Sprintf("%d-%d-%d %d:%d:%d %s: %s\n", t.Day(), t.Month(), t.Year(), t.Hour(), t.Minute(), t.Second(), level, msg)
		fmt.Println(strMsg)
		n, err := fp.WriteString(strMsg)
		if err != nil{
			panic(err)
		}
		fmt.Println(n)
	}
}

func Info(msg string){
	logPrint(msg, Info_msg)
}

func Warn(msg string){
	logPrint(msg, Warn_msg)
}

func Error(msg string){
	logPrint(msg, Error_msg)
}

func Fatal(msg string){
	logPrint(msg, Fatal_msg)
}

/*func main(){
	Init("sample.log")
	//log.logPrint("test message", "INFO")
	Info("Info message")
	Warn("Warn message")
	Error("Error message")
	Fatal("fatal message")
}*/
