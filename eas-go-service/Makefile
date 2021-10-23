# eg. 编译所有服务：make
# eg. 编译某个服务：make eas-admin

all: clean init build

clean:
	rm -rf target

init:
	mkdir -p "target"

build:
	bash build.sh binary $(service)

#-----------------------------------------------------------------------------------------------
docker: docker-build docker-push docker-clean

docker-build:
	bash build.sh docker-build $(service)

docker-push:
	bash build.sh docker-push $(service)

docker-clean:
	bash build.sh docker-clean $(service)
