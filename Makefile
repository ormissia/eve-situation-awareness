# eg. 编译所有服务：make
# eg. 编译某个服务：make eas-admin

all: clean init build

clean:
	rm -rf target

init:
	mkdir -p "target"

build:
	cd eas-go-service && make
