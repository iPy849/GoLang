package main

// TODO: Necesito instalar XCode para tener acceso al logger
// TODO: Revisar de nuevo la sección de Logging Information en el capítulo 1

/*
xcode-select --install
https://www.cyberithub.com/solved-xcrun-error-invalid-active-developer-path-library-develop/
*/

import (
	"fmt"
	"log"
	"log/syslog"
)

func main() {
	sysLog, err := syslog.New(syslog.LOG_SYSLOG, "systemLog.log")
	if err != nil {
		fmt.Println(err)
		return
	} else {
		log.SetOutput(sysLog)
		log.Println("Everithing is fine")
	}
}
