package main

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

const dateForm = "Jan 2 15:04:05 2006"

var (
	// zpoolTotalDisks = prometheus.NewDesc(
	// 	"zfs_zpool_total_disks",
	// 	"Total number of disks in zpool",
	// 	[]string{"name"},
	// 	nil,
	// )
	zpoolError = prometheus.NewDesc(
		"zfs_zpool_error",
		"Is there a zpool error",
		[]string{"name"},
		nil,
	)
	zpoolOnline = prometheus.NewDesc(
		"zfs_zpool_online",
		"Is the zpool online",
		[]string{"name"},
		nil,
	)
	zpoolLastScrub = prometheus.NewDesc(
		"zfs_zpool_last_scrub",
		"Last zpool scrub",
		[]string{"name"},
		nil,
	)
	zpoolLastScrubErrors = prometheus.NewDesc(
		"zfs_zpool_last_scrub_errors",
		"Last scrub total errors on the zpool",
		[]string{"name"},
		nil,
	)
	zpoolParsingError = prometheus.NewDesc(
		"zfs_zpool_parsing_error",
		"Error when trying to parse the API data.",
		[]string{"name"},
		nil,
	)
)

type proxmoxZpoolCollector struct {
	name string
	api  *ProxmoxAPI
}

func newProxmoxZpoolCollector(name string, api *ProxmoxAPI) *proxmoxZpoolCollector {
	return &proxmoxZpoolCollector{
		name: name,
		api:  api,
	}
}

func (collector *proxmoxZpoolCollector) Describe(ch chan<- *prometheus.Desc) {
	//ch <- zpoolTotalDisks
	ch <- zpoolError
	ch <- zpoolOnline
	ch <- zpoolLastScrub
	ch <- zpoolLastScrubErrors
	ch <- zpoolParsingError
}

func (collector *proxmoxZpoolCollector) Collect(ch chan<- prometheus.Metric) {
	zpoolList, err := collector.api.GetZpoolList()
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, zpool := range zpoolList.Data {
		var zpoolParsingErrorMetric float64

		zpoolInfo, err := collector.api.GetZpool(zpool.Name)
		if err != nil {
			zpoolParsingErrorMetric = float64(1)
			fmt.Println(err)
		}

		var zpoolOnlineMetric float64
		var zpoolErrorMetric float64
		if zpoolInfo.Data.State == "ONLINE" {
			zpoolOnlineMetric = float64(1)
			zpoolErrorMetric = float64(0)
		} else {
			zpoolErrorMetric = float64(1)
		}

		var zpoolLastScrubMetric float64
		//Example scrub response: scrub repaired 0B in 0 days 01:56:29 with 0 errors on Sun May 10 02:20:30 2020
		if x := strings.SplitAfter(zpoolInfo.Data.Scan, "on "); len(x) == 2 {
			//Sun May 10 02:20:30 2020
			if len(x[1]) > 5 { //We want to get rid of the day eg: Mon
				//May 10 02:20:30 2020
				t, err := time.Parse(dateForm, x[1][4:])
				if err != nil {
					zpoolParsingErrorMetric = float64(1)
					fmt.Println(err)
				}
				zpoolLastScrubMetric = float64(t.Unix()) //Could this be an issue since time.Unix() returns int64?
			}
		}

		var zpoolLastScrubErrorsMetric float64
		//Example scrub response: scrub repaired 0B in 0 days 01:56:29 with 0 errors on Sun May 10 02:20:30 2020
		splitLine := strings.Split(zpoolInfo.Data.Scan, " ")
		for index, x := range splitLine {
			fmt.Println(index)
			if strings.Contains(x, "error") && index >= 1 { //Support for "error" or "errors"
				totalErrors, err := strconv.ParseFloat(splitLine[index-1], 64) //We want to grab the number before error eg: 3 errors
				if err != nil {
					zpoolParsingErrorMetric = float64(1)
				} else {
					zpoolLastScrubErrorsMetric = totalErrors
					fmt.Println(zpoolLastScrubErrorsMetric)
				}
				break
			}
		}

		//ch <- prometheus.MustNewConstMetric(zpoolTotalDisks, prometheus.GaugeValue, metricValue, "test")
		ch <- prometheus.MustNewConstMetric(zpoolError, prometheus.GaugeValue, zpoolErrorMetric, zpool.Name)
		ch <- prometheus.MustNewConstMetric(zpoolOnline, prometheus.GaugeValue, zpoolOnlineMetric, zpool.Name)
		ch <- prometheus.MustNewConstMetric(zpoolLastScrub, prometheus.GaugeValue, zpoolLastScrubMetric, zpool.Name)
		ch <- prometheus.MustNewConstMetric(zpoolLastScrubErrors, prometheus.GaugeValue, zpoolLastScrubErrorsMetric, zpool.Name)
		ch <- prometheus.MustNewConstMetric(zpoolParsingError, prometheus.GaugeValue, zpoolParsingErrorMetric, zpool.Name)
	}
}
