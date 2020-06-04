package main

import (
	"log"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var officeHours = prometheus.NewGauge(prometheus.GaugeOpts{
	Name:        "office_hours",
	Help:        "Determines whether current time is during office hours.",
	ConstLabels: prometheus.Labels{"location": "Toronto"},
})

var wakingHours = prometheus.NewGauge(prometheus.GaugeOpts{
	Name:        "waking_hours",
	Help:        "Determines whether current time is during waking hours.",
	ConstLabels: prometheus.Labels{"location": "Toronto"},
})

func init() {
	prometheus.MustRegister(officeHours)
	prometheus.MustRegister(wakingHours)
}

func pollingTime() {
	loc, err := time.LoadLocation("America/Toronto")
	if err != nil {
		log.Fatal(err)
	}

	for {
		t := time.Now().In(loc)

		officeHours.Set(BoolToFloat64(duringOfficeHours(t)))
		wakingHours.Set(BoolToFloat64(duringWakingHours(t)))

		time.Sleep(300 * time.Second)
	}
}

func main() {
	go pollingTime()

	http.Handle("/metrics", promhttp.Handler())

	log.Print("Metrics being served on port: 9001")
	log.Fatal(http.ListenAndServe(":9001", nil))
}

func duringOfficeHours(t time.Time) bool {
	// Office hours: Monday - Friday, 9:00AM until 5:00PM
	return t.Hour() >= 9 && t.Hour() < 17 && t.Weekday() != time.Saturday && t.Weekday() != time.Sunday
}

func duringWakingHours(t time.Time) bool {
	// Waking hours: Monday - Friday, 7:00AM until 11:00PM
	return t.Hour() >= 8 && t.Hour() < 23 && t.Weekday() != time.Saturday && t.Weekday() != time.Sunday
}
