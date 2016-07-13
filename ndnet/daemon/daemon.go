package daemon

import (
	log "github.com/Sirupsen/logrus"
	dn "github.com/docker/go-plugins-helpers/network"
)

var (
	PLUGIN_NAME = "ndnet"
	PLUGIN_PORT = ":2804"
)

func Start(cfgFile string, debug bool) {
	if debug == true {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	d, err := DriverAlloc(cfgFile)
	if err != nil {
		log.Fatalf("ERROR: %s init failed!", PLUGIN_NAME)
	}
	h := dn.NewHandler(d)
	log.Info("%s driver created & initialized", PLUGIN_NAME)
	log.Info(h.ServeTCP(PLUGIN_NAME, PLUGIN_PORT))
}
