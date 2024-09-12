module hoymiles_wifi

go 1.21

toolchain go1.23.0

require (
	github.com/basvdlei/gotsmart v0.0.3
	google.golang.org/protobuf v1.34.2
)

require github.com/google/go-cmp v0.6.0 // indirect

replace github.com/grid-x/modbus => github.com/evcc-io/modbus v0.0.0-20240503125516-9fd99fe0e438
