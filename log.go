package logging

import (
	"io"
	"log"
	"os"
	
	"github.com/sec51/goconf"

)

var (
	Trace   *log.Logger
	Info    *log.Logger
	Warning *log.Logger
	Error   *log.Logger
)

func init() {

	logFile := goconf.AppConf.String("logging::file")

	var multi io.Writer
	var file *os.File
	var err error
	if logFile != "" {
		file, err = os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file %s with error: %s\n", logFile, err)
		}
	}

	if logFile != "" {
		multi = io.MultiWriter(file, os.Stdout)
	} else {
		multi = io.MultiWriter(os.Stdout)
	}

	Trace = log.New(multi,
		"TRACE: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Info = log.New(multi,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Warning = log.New(multi,
		"WARNING: ",
		log.Ldate|log.Ltime|log.Lshortfile)

	Error = log.New(multi,
		"ERROR: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

