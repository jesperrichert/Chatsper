rm -rf ./internal/api/router/rice-box.go
cd ./internal/api/router/ && rice embed-go
cd ../../../ && go build ./cmd/main.go