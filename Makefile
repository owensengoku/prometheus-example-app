VERSION=v0.1.0

all:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o prometheus-example-app --installsuffix cgo main.go
	docker build -t quay.io/owensengoku/prometheus-example-app:$(VERSION) .
