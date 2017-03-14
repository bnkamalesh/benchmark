package main

import (
	"errors"
	"io/ioutil"
	"net"
	"net/http"
	"time"

	"github.com/bnkamalesh/benchmark"
)

func dialTimeout(network, addr string) (net.Conn, error) {
	return net.DialTimeout(network, addr, time.Duration(15*time.Second))
}

var httpClient http.Client

// The benchFn interface implementation
type benchFn struct{}

func (bF benchFn) Call() error {
	req, _ := http.NewRequest("GET", "https://kamaleshwar.com", nil)

	resp, err := httpClient.Do(req)
	if err != nil {
		return err
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()
	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		errResp := string(body)
		return errors.New(errResp)
	}

	return err
}

func main() {
	httpClient = http.Client{
		Timeout: time.Second * time.Duration(10),
		Transport: &http.Transport{
			Dial: dialTimeout,
		},
	}

	// Returns a new Benchmark pointer with all the defaults assigned
	benchmark := benchmark.New()
	// time to wait before firing the consequent request
	benchmark.WaitPerReq = time.Millisecond * 1
	// print available stats while the benchmark is running
	benchmark.ShowProgress = true
	// Total number of requests to fire
	benchmark.TotalRequests = 100
	// Duration in which all the requests have to be finished firing (in milliseconds).
	benchmark.BenchDuration = 1000

	// Updates all the necessary fields according to the configuration provided
	benchmark.Init()
	// Run the benchmark
	benchmark.Run(benchFn{})
}
