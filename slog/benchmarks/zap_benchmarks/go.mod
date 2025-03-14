module github.com/jcbhmr/exp/slog/benchmarks/zap_benchmarks

go 1.23.0

require (
	go.uber.org/zap v1.24.0
	github.com/jcbhmr/exp v0.0.0-20220909182711-5c715a9e8561
)

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

replace github.com/jcbhmr/exp => ../../..
