all: clean init build

clean:
	rm -rf target

init:
	mkdir -p "target"

build: eas-admin-service eas-source-zkill-service

#TODO
#-------------------
docker: docker-build docker-push

docker-build: docker-build-eas-admin docker-build-eas-source-zkill

docker-push:

#-------------------
eas-admin-service:
	GOOS=linux GOARCH=amd64 go build -o target/eas-admin eas-admin/main.go

docker-build-eas-admin:
	docker build -t ormissia/eas-admin --target eas-admin .

eas-source-zkill-service:
	GOOS=linux GOARCH=amd64 go build -o target/eas-source-zkill eas-source-zkill/main.go

docker-build-eas-source-zkill:
	docker build -t ormissia/eas-source-zkill --target eas-source-zkill .