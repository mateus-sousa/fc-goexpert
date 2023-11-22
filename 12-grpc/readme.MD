Antes, precisa do protoc instalado:
https://grpc.io/docs/protoc-installation/

Gera entidades go:
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28

Gera interfaces de comunicação grpc:
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2

dps de criar o arquivo .proto, gere as interafes grpc:
protoc --go_out=. --go-grpc_out=. proto/course_category.proto
go mod tidy

programa para interagir com grpc:
https://github.com/ktr0731/evans

para utilizar o evans:
evans -r repl

selecione o package, selecione o service e chame a função.