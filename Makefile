
protoc:
	protoc --proto_path=protos --go_out=protos --go_opt=paths=source_relative protos/*.proto