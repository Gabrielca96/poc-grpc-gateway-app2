plugins
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

protoc --go_out=./gen --go-grpc_out=./gen --proto_path=proto proto/*.proto


go mod tidy


aws ecr create-repository \
    --repository-name odihnx-poc-grpc-app-b \
    --region us-east-1 \
    --profile odihnx-dev



aws ecr get-login-password --region us-east-1 --profile odihnx-dev | \
docker login --username AWS --password-stdin 590184058323.dkr.ecr.us-east-1.amazonaws.com


docker build -t poc-grpc-app-b .

docker tag poc-grpc-app-b:latest 590184058323.dkr.ecr.us-east-1.amazonaws.com/odihnx-poc-grpc-app-b:latest

docker push 590184058323.dkr.ecr.us-east-1.amazonaws.com/odihnx-poc-grpc-app-b:latest