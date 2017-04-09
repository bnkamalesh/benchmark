# benchmark
A Go library to get benchmarks of a function (execution time). All requests/executions are done in *parellel* (Go routines).

## How to use?
Any function with the signature `fn() error` can be used to run the benchmark. Please refer to the sample code `/sample/main.go`


## Sample output:

Benchmark run for HTTP GET request to `https://kamaleshwar.com`.

```
Duration: 1s  Total requests: 100  Wait time per request: 10ms  Show progess: true , per 10 request(s)

Start:2017-03-14 20:48:02.432076782 +0530 IST

10  out of  100  done.  Success: 9  Errors: 0
20  out of  100  done.  Success: 19  Errors: 0
30  out of  100  done.  Success: 29  Errors: 0
40  out of  100  done.  Success: 39  Errors: 0
50  out of  100  done.  Success: 49  Errors: 0
60  out of  100  done.  Success: 59  Errors: 0
70  out of  100  done.  Success: 69  Errors: 0
80  out of  100  done.  Success: 79  Errors: 0
90  out of  100  done.  Success: 89  Errors: 0
100  out of  100  done.  Success: 94  Errors: 4

Done: 2017-03-14 20:48:12.438243541 +0530 IST 

========================= Benchmark stats =========================
 
Total requests: 100 
Requests completed: 100
------

Success: 94 (94%) 
Errors: 6 (6%)
------

Total time taken to complete the benchmark: 10.006308073s
Average time per successful request: 2.335882319s 
Fastest: 323.045732ms 
Slowest: 8.266474823s

Average time per failed request: 10.004783269s 
Fastest: 10.00394276s 
Slowest: 10.005461733s
------

Error messages (1)

 1. Get https://kamaleshwar.com: net/http: request canceled (Client.Timeout exceeded while awaiting headers)
  Occurrences: 4
```
