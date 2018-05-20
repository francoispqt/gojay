# HTTP BENCHMARKS

This package has two different implementation of a basic HTTP server in Go. It just takes a JSON body, unmarshals it and marshals it back to the response. 

It required `wrk` benchmarking tool, which you can find here: https://github.com/wg/wrk

## How to run 

### gojay
```bash
cd /path/to/package && go run gojay/main.go
```
Then run: 
```bash
cd /path/to/package && make bench
```

### standard package (encoding/json)
```bash
cd /path/to/package && go run standard/main.go
```
Then run: 
```bash
cd /path/to/package && make bench
```

## Results 

Results presented here are ran on a MacBook Pro Mid 2015 2,2 GHz Intel Core i7 with 16G of RAM

**gojay results:**
```
Running 20s test @ http://localhost:3000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   298.77us  341.40us  10.52ms   94.13%
    Req/Sec    18.88k     1.89k   21.40k    73.63%
  755246 requests in 20.10s, 1.67GB read
Requests/sec:  37573.84
Transfer/sec:     84.85MB
```
**standard results:**
```
Running 20s test @ http://localhost:3000
  2 threads and 10 connections
  Thread Stats   Avg      Stdev     Max   +/- Stdev
    Latency   613.21us  557.50us  12.65ms   93.88%
    Req/Sec     9.18k   423.20    10.10k    80.50%
  365404 requests in 20.00s, 811.60MB read
Requests/sec:  18269.66
Transfer/sec:     40.58MB
```