genG:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative calculator/calculator_pb/*.proto
clean:
	rm calculator/calculator_pb/*.go
runs:
	go run calculator/calculator_server/server.go
runc:
	go run calculator/calculator_client/client.go