.PHONY: test
test:
	go test -run=^Test -v

.PHONY: bench
bench: 
	go test -benchmem -run=^$ github.com/francoispqt/gojay/benchmarks/decoder -bench ^Benchmark

.PHONY: benchtrace
benchtrace:
	go test -c & GODEBUG=allocfreetrace=1 ./decoder.test -test.run=none -test.bench=^Benchmark -test.benchtime=10ms 2>trace.log

.PHONY: benchcpu
benchcpu: 
	go test -benchmem -run=^$ github.com/francoispqt/gojay/benchmarks/decoder -bench ^Benchmark -cpuprofile cpu.out

.PHONY: testtrace
testtrace: 
	go test -benchmem -run=^$ github.com/francoispqt/gojay/benchmarks/decoder -bench ^Benchmark -trace trace.out

.PHONY: benchgojay
benchgojay:
	go test -benchmem -run=^BenchmarkGoJay -bench=^BenchmarkGoJay -benchtime=30ms

.PHONY: benchgojaycpu
benchgojaycpu: 
	go test -benchmem -run=^BenchmarkGoJay -bench=^BenchmarkGoJay -benchtime=30ms -cpuprofile cpu.out

.PHONY: benchjsoniter
benchjsoniter:
	go test -benchmem -run=^BenchmarkJsonIter -bench=^BenchmarkJsonIter -benchtime=30ms

.PHONY: benchjsonparser
benchjsonparser:
	go test -benchmem -run=^BenchmarkJsonParser -bench=^BenchmarkJsonParser -benchtime=30ms