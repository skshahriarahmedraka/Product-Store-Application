update:
	go get -u
	go mod tidy
	# go mod vendor
	go mod verify

test:
	go test -v ./...
	# go test `go list ./... | grep -v e2e_test`
run:
	go run main.go
build:
	go build -o main main.go

docker_test: docker-build_test 
	sudo docker compose up -d 
	sudo docker compsoe exec -T http go test ./...
	sudo docker compose down 
docker_build:
	sudo docker build -t shahriarraka/product_store_application:latest .
	# docker build -t shahriarraka/product_store_application:latest .
docker_push:
	sudo docker push shahriarraka/product_store_application:latest 
docker_build_test:
	docker build . -t service_test --target=test

docker_run:
	sudo docker stop product_store_application1 || true && sudo docker rm product_store_application1 || true
	sudo docker run --publish 8080:8080 --name product_store_application1 shahriarraka/product_store_application:latest  