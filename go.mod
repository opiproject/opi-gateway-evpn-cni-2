module github.com/opiproject/opi-gateway-evpn-cni

go 1.20

require (
	github.com/containernetworking/cni v1.1.2
	github.com/containernetworking/plugins v1.2.0
	github.com/imdario/mergo v0.3.15
	github.com/k8snetworkplumbingwg/sriovnet v1.2.0
	github.com/opiproject/godpu v0.2.1-0.20240412165547-e4af793befd9
	github.com/opiproject/opi-api v0.0.0-20240304222410-5dba226aaa9e
	github.com/opiproject/opi-evpn-bridge v0.1.1-0.20240425152645-d33fbefc0eb4
	github.com/vishvananda/netlink v1.2.1-beta.2.0.20240226175043-124bb8e72178
	golang.org/x/net v0.21.0
	google.golang.org/grpc v1.61.0
)

require (
	github.com/cenkalti/backoff/v4 v4.2.1 // indirect
	github.com/coreos/go-iptables v0.6.0 // indirect
	github.com/go-logr/logr v1.4.1 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/uuid v1.5.0 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.19.0 // indirect
	github.com/safchain/ethtool v0.2.0 // indirect
	github.com/spf13/afero v1.9.4 // indirect
	github.com/vishvananda/netns v0.0.4 // indirect
	github.com/ziutek/telnet v0.0.0-20180329124119-c3b780dc415b // indirect
	go.einride.tech/aip v0.66.0 // indirect
	go.opentelemetry.io/otel v1.24.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.21.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.21.0 // indirect
	go.opentelemetry.io/otel/metric v1.24.0 // indirect
	go.opentelemetry.io/otel/sdk v1.21.0 // indirect
	go.opentelemetry.io/otel/trace v1.24.0 // indirect
	go.opentelemetry.io/proto/otlp v1.0.0 // indirect
	golang.org/x/sys v0.17.0 // indirect
	golang.org/x/text v0.14.0 // indirect
	google.golang.org/genproto v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240108191215-35c7eff3a6b1 // indirect
	google.golang.org/protobuf v1.32.0 // indirect
)

replace github.com/opiproject/opi-evpn-bridge => github.com/mardim91/opi-evpn-bridge v0.0.0-20240426100730-82abad652a1e
