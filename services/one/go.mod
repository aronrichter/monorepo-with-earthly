module github.com/aronrichter/monorepo-with-earthly/services/one

go 1.22.3

require (
	github.com/aronrichter/monorepo-with-earthly/libs/log v0.0.0
	go.uber.org/zap v1.27.0
)

replace github.com/aronrichter/monorepo-with-earthly/libs/log v0.0.0 => ../../libs/log

require go.uber.org/multierr v1.11.0 // indirect
