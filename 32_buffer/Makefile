bench:
	go test -bench=BenchmarkStart -benchmem -benchtime=5s -count=5

pprof:
	go test -bench=BenchmarkStart -benchmem -benchtime=5s -cpuprofile cpu.out -memprofile mem.out

pprof-mem:
	go tool pprof -http :9000 mem.out

pprof-cpu:
	go tool pprof -http :9000 cpu.out



