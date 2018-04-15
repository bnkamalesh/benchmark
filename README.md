[![Build Status](https://travis-ci.org/bnkamalesh/benchmark.svg?branch=master)](https://travis-ci.org/bnkamalesh/benchmark)
[![](https://goreportcard.com/badge/github.com/bnkamalesh/benchmark)](https://goreportcard.com/report/github.com/bnkamalesh/benchmark)
[![](https://cover.run/go/github.com/bnkamalesh/benchmark.svg?tag=golang-1.10)](https://cover.run/go/github.com/bnkamalesh/benchmark)
[![](https://godoc.org/github.com/nathany/looper?status.svg)](http://godoc.org/github.com/bnkamalesh/benchmark)

# Benchmark
A Go library to get benchmarks of a function (execution time). All requests/executions are done in *parallel* (Go routines).

## How to use?
Any function with the signature `fn() error` can be used to run the benchmark. Please refer to the Example provided in the test file.


## Sample output:

Benchmark run for HTTP GET request to `https://kamaleshwar.com`.

```
Duration              : 1s 
Total requests        : 750 
Wait time per request : 1.333333ms 
Show progess          : true , per 75 request(s) 
Start                 : 2017-04-25 03:56:18.917402373 +0530 IST


75  out of  750  done.  Success: 74  Errors: 0
150  out of  750  done.  Success: 149  Errors: 0
225  out of  750  done.  Success: 224  Errors: 0
300  out of  750  done.  Success: 299  Errors: 0
375  out of  750  done.  Success: 374  Errors: 0
450  out of  750  done.  Success: 449  Errors: 0
525  out of  750  done.  Success: 524  Errors: 0
600  out of  750  done.  Success: 599  Errors: 0
675  out of  750  done.  Success: 674  Errors: 0
750  out of  750  done.  Success: 748  Errors: 0

========================= Benchmark stats =========================
 
Done               : 2017-04-25 03:56:29.019980966 +0530 IST 
Time to complete   : 10.1025787s 
Total requests     : 750 
Requests completed : 750 
Success            : 748 (99.73333333333333%) 
Errors             : 2 (0.26666666666666666%)

Average time per successful request : 3.291541039s 
Fastest                             : 256.852232ms 
Slowest                             : 7.255390376s

Average time per failed request : 10.004558347s 
Fastest                         : 10.003726326s 
Slowest                         : 10.005390368s


Error messages (1)

 1. Get https://kamaleshwar.com: net/http: request canceled (Client.Timeout exceeded while awaiting headers)
  Occurrences: 2
```
