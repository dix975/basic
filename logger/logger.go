package logger

import (
	"io"
	"log"
	"os"
	"io/ioutil"
)



var (
	Trace   *log.Logger = log.New(ioutil.Discard, "TRACE: ", log.Ldate | log.Ltime | log.Lshortfile)
	Debug   *log.Logger = log.New(ioutil.Discard, "TRACE: ", log.Ldate | log.Ltime | log.Lshortfile)
	Info    *log.Logger = log.New(ioutil.Discard, "TRACE: ", log.Ldate | log.Ltime | log.Lshortfile)
	Warning *log.Logger = log.New(ioutil.Discard, "TRACE: ", log.Ldate | log.Ltime | log.Lshortfile)
	Error   *log.Logger = log.New(ioutil.Discard, "TRACE: ", log.Ldate | log.Ltime | log.Lshortfile)
)

func Init() {


	InitWithIO(os.Stdout, os.Stdout, os.Stdout, os.Stderr)
}

func InitWithIO(traceHandle io.Writer, infoHandle io.Writer, warningHandle io.Writer, errorHandle io.Writer) {

	//Trace = log.New(traceHandle, "TRACE: ", log.Ldate | log.Ltime | log.Lshortfile)
	Debug = log.New(traceHandle, "DEBUG: ", log.Ldate | log.Ltime | log.Lshortfile)
	Info = log.New(infoHandle, "INFO: ", log.Ldate | log.Ltime | log.Lshortfile)
	Warning = log.New(warningHandle, "WARNING: ", log.Ldate | log.Ltime | log.Lshortfile)
	Error = log.New(errorHandle, "ERROR: ", log.Ldate | log.Ltime | log.Lshortfile)

}

//todo : finish the implementation (http://www.goinggo.net/2013/11/using-log-package-in-go.html)