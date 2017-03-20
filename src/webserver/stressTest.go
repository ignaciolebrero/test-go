package main

import (
  "fmt"
  "time"

  vegeta "github.com/tsenart/vegeta/lib"
)

func main() {
  rate := uint64(100) // per second
  duration := 4 * time.Second
  targeter := vegeta.NewStaticTargeter(vegeta.Target{
    Method: "GET",
    URL:    "http://localhost:8080/prueba",
  })
  attacker := vegeta.NewAttacker()

  var metrics vegeta.Metrics
  for res := range attacker.Attack(targeter, rate, duration) {
    metrics.Add(res)
  }
  metrics.Close()

  fmt.Printf("latencies.mean: %s\n", metrics.Latencies.Mean)
  fmt.Printf("latencies.max: %s\n", metrics.Latencies.Max)
  // fmt.Printf("requests.total: %s\n", metrics.Requests.Total)
  fmt.Printf("Requests.rate: %s\n", metrics.Requests)
}
