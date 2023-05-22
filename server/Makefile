.PHONY: run build image push clean

tag = v0.1
releaseName = lbemi-server
image = registry.cn-hangzhou.aliyuncs.com/lbemi/crmdb

ALL: run

run: build
	./lbemi-server

build:
	go build -o $(releaseName) ./cmd/

image:
	docker build -t ${image}:$(tag) .

push: image
	docker push  ${image}:$(tag)

clean:
	-rm -f ./$(releaseName)

