package logging

import (
	"io"
	"log"
)

var Info *log.Logger
var Warn *log.Logger

const logPattern = log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile | log.LUTC

func InitLogs(infoHandle io.Writer, warnHandle io.Writer, panicHandle io.Writer) {
	//to be used for INFO-level logging: Info.Println("foo is now infobar")
	Info = log.New(infoHandle, "INFO  - ", logPattern)
	//to be used for WARN-level logging: Warn.Println("foo is now warnbar")
	Warn = log.New(warnHandle, "WARN  - ", logPattern)

	//log is used for ERROR-level logging: log.Println("foobar error")
	log.SetFlags(logPattern)
	log.SetPrefix("ERROR - ")
	log.SetOutput(panicHandle)
}
