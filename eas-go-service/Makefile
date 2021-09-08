SERVICE_NAME = N

all: clean init build

clean:
	rm -rf target

init:
	mkdir -p "target"

build:
	sh build.sh $(sname)

#-----------------------------------------------------------------------------------------------
docker: docker-build docker-push docker-clean

docker-build: eas-admin-docker eas-source-zkill-docker

docker-push:

docker-clean:

eas-admin-docker:
	docker build -t ormissia/eas-admin --target eas-admin .

#-----------------------------------------------------------------------------------------------
eas-source-zkill-docker:
	docker build -t ormissia/eas-source-zkill --target eas-source-zkill .