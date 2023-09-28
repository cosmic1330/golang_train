```cmd
brew install protobuf
brew install protoc-gen-go(macOS)
brew install protoc-gen-go-grpc(macOS)

go install google.golang.org/protobuf/cmd/protoc-gen-go@latest

go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```
接著一樣是終端機，我們切到專案內 proto 的資料夾底下後，輸入

```cmd
protoc --go_out=. --go-grpc_out=. test.proto
```

缺點
不能透過瀏覽器直接調用gRPC服務。
gRPC是使用Protobuf進行編碼，雖可高效率傳送，但人工不可讀。
使用時須先定義Protobuf規範，雖嚴謹但較費時。