echo "Running Unit Tests..."
go test ./pkg/engine/...

echo "Running Benchmarks..."
go test -bench=. ./tests/bench_test.go