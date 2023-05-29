service-transcriber-api

A repository for the  transcriber service api being developed 
for ant investors to convert all their speeches into text

### How do I update the definitions? ###

* The api definition is defined in the proto file transcriber.proto
* To update the proto service you need to run the command :


    `protoc --proto_path=../apis --proto_path=./v1 --go_out=./ --validate_out=lang=go:. transcriber.proto`

    `protoc --proto_path=../apis --proto_path=./v1  transcriber.proto --go-grpc_out=./ `
    
    `mockgen -source=notification_grpc.pb.go -self_package=github.com/antinvestor/service-transcriber-api -package=transcriberv1 -destination=transcriber_grpc_mock.go`

with that in place update the implementation appropriately
