package configs

import (
	"os"

	log "github.com/sirupsen/logrus"
	prefixed "github.com/t-tomalak/logrus-prefixed-formatter"
)

// SetupLogger initializes global logger
func SetupLogger() {

	log.SetOutput(os.Stdout)
	log.SetFormatter(&prefixed.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
		ForceFormatting: true,
	})
	log.SetLevel(log.DebugLevel)

}
