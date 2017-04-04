package main
import "./myLogger"

func main(){
	myLogger.Init("test.log")
	myLogger.Info("testing logger")
	myLogger.Warn("testing logger")
	myLogger.Error("testing logger")
}
