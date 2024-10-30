proto:
	protoc --go_out=. --go-grpc_out=. *.proto
	protoc pkg/**/pb/*.proto --go_out=plugins=grpc:.
	protoc --go_out=. --go-grpc_out=. auth.proto
	protoc --go-grpc_out=. *.proto
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative auth/auth.proto

