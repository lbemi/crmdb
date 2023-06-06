.PHONY: run build image push clean

tag = v1.0
releaseName = crmdb
image = registry.cn-hangzhou.aliyuncs.com/lbemi/crmdb

ALL: run

run: build
	./crmdb

build:
	go build -o $(releaseName) ./cmd/

image:
	docker build -t ${image}:$(tag) .

push: image
	docker push  ${image}:$(tag)

clean:
	-rm -f ./$(releaseName)

