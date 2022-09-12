module validator

go 1.18

replace customer_service_proto => ../customer_service_proto

require (
	customer_service_proto v0.0.0-00010101000000-000000000000
	github.com/segmentio/kafka-go v0.4.35
	google.golang.org/grpc v1.49.0
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/klauspost/compress v1.15.7 // indirect
	github.com/pierrec/lz4/v4 v4.1.15 // indirect
	golang.org/x/net v0.0.0-20220706163947-c90051bbdb60 // indirect
	golang.org/x/sys v0.0.0-20220520151302-bc2c85ada10a // indirect
	golang.org/x/text v0.3.7 // indirect
	google.golang.org/genproto v0.0.0-20200526211855-cb27e3aa2013 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)
