package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"math/rand"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"
	"net/http"
	"os"
)

type ClusterManager struct {
	Zone string
	OOM *prometheus.Desc
	Ram *prometheus.Desc
}

func (c *ClusterManager) C1()(oom map[string]int, ram map[string]float64)  {

	oom = map[string]int{
		"test1":int(rand.Int31n(1000)),
		"test2":int(rand.Int31n(1000)),
	}

	ram = map[string]float64{
		"test1":rand.Float64() * 100,
		"test2":rand.Float64() * 100,
	}

	return
}

func (c *ClusterManager) Describe(ch chan <- *prometheus.Desc)  {

	ch <- c.OOM
	ch <- c.Ram
}

func (c *ClusterManager) Collect(ch chan <- prometheus.Metric)  {
	oom,ram := c.C1()
	for h,oc := range oom {
		ch <- prometheus.MustNewConstMetric(c.OOM,prometheus.CounterValue,float64(oc),h)
	}
	for h,r := range ram {
		ch <- prometheus.MustNewConstMetric(c.OOM,prometheus.GaugeValue,r,h)
	}
}

func NewClusterManager(zone string) *ClusterManager  {
	return &ClusterManager{
		Zone:zone,
		OOM:prometheus.NewDesc("crash_total","total cresh count",[]string{"host"},prometheus.Labels{"zone":zone}),
		Ram:prometheus.NewDesc("ram_usage","ram warning",		[]string{"host"},	prometheus.Labels{"zone":zone}),
	}
}

func main() {

	//http.Handle("/metrics",promhttp.Handler())
	//log.Fatal(http.ListenAndServe(":8080",nil))

	workerA := NewClusterManager("A")
	workerB := NewClusterManager("B")
	reg := prometheus.NewPedanticRegistry()
	reg.Register(workerA)
	reg.Register(workerB)

	gatherers := prometheus.Gatherers{
		prometheus.DefaultGatherer,reg,
	}

	h := promhttp.HandlerFor(gatherers,promhttp.HandlerOpts{
		ErrorLog:log.NewErrorLogger(),
		ErrorHandling:promhttp.ContinueOnError,
	})
	http.HandleFunc("/metrics", func(writer http.ResponseWriter, request *http.Request) {
		h.ServeHTTP(writer,request)
		
	})

	log.Infoln("start server at :8080")
	if err := http.ListenAndServe(":8080",nil); err != nil {
		log.Errorf("Error occured at server start %v",err)
		os.Exit(-1)
	}
}
