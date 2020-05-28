package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func readConfigFile() ProxmoxAPI {
	config := ProxmoxAPI{}
	f, err := os.Open("/etc/proxmox-zfs-exporter/config.json")
	if err != nil {
		panic("Cannot open file.")
	}
	defer f.Close()

	enc := json.NewDecoder(f)
	err = enc.Decode(&config)
	if err != nil {
		panic("Cannot decode config file.")
	}

	return config
}

func main() {
	proxmoxAPI := readConfigFile()
	collector := newProxmoxZpoolCollector("test", &proxmoxAPI)
	prometheus.MustRegister(collector)

	go proxmoxAPI.refreshTicket()
	//Wait for the first ticket to be set
	proxmoxAPI.waitForTicket()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9000", nil))
}
