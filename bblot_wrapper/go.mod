module github.com/DaiYuANg/gokit/bblot_wrapper

go 1.25

require (
	github.com/DaiYuANg/gokit/codec v0.0.0-00010101000000-000000000000
	github.com/DaiYuANg/gokit/codec/json_codec v0.0.0-00010101000000-000000000000
	github.com/samber/lo v1.52.0
	github.com/samber/oops v1.19.4
	go.etcd.io/bbolt v1.4.3
)

require (
	github.com/oklog/ulid/v2 v2.1.1 // indirect
	go.opentelemetry.io/otel v1.29.0 // indirect
	go.opentelemetry.io/otel/trace v1.29.0 // indirect
	golang.org/x/sys v0.29.0 // indirect
	golang.org/x/text v0.22.0 // indirect
)

replace github.com/DaiYuANg/gokit/codec => ../codec

replace github.com/DaiYuANg/gokit/codec/json_codec => ../codec/json_codec
