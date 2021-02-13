package logger

import (
	"os"

	"github.com/NodeHodl/btc-node-proxy/config"

	logrus "github.com/sirupsen/logrus"
)

//New ...
func New(module string) *logrus.Entry {
	log := logrus.New()

	//If you wish to add the calling method as a field, instruct the logger via:
	// log.SetReportCaller(true)

	// Log as JSON instead of the default ASCII formatter.
	// log.SetFormatter(&logrus.JSONFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)

	// Only log the warning severity or above.
	// log.SetLevel(logrus.WarnLevel)

	// A common pattern is to re-use fields between logging statements by re-using
	// the logrus.Entry returned from WithFields()
	contextLogger := log.WithFields(logrus.Fields{
		"service": config.ServiceName,
		"module":  module,
	})

	return contextLogger
}
