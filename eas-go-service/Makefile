all: clean init build

clean:
	rm -rf target

init:
	mkdir -p "target"

build: eas-admin-service eas-source-zkill-service

#TODO
#-----------------------------------------------------------------------------------------------
docker: docker-build docker-push docker-clean

docker-build: eas-admin-docker eas-source-zkill-docker

docker-push:

docker-clean:

#-----------------------------------------------------------------------------------------------
eas-admin-service:
	go build -o target/eas-admin eas-admin/main.go

eas-admin-docker:
	docker build -t ormissia/eas-admin --target eas-admin .

#-----------------------------------------------------------------------------------------------
eas-source-zkill-service:
	go build -o target/eas-source-zkill eas-source-zkill/main.go

eas-source-zkill-docker:
	docker build -t ormissia/eas-source-zkill --target eas-source-zkill .