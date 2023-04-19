.PHONY: run build image push clean

tag = v0.1
releaseName = crmdb
image = registry.cn-hangzhou.aliyuncs.com/lbemi/crmdb
dockerhubUser = lei@1996262755539687

ALL: run

run: build
	./crmdb

build:
	go build -o $(releaseName) ./cmd/

image:
	docker build -t ${image}:$(tag) .

push: image
	docker push $(dockerhubUser)/$(releaseName):$(tag)

clean:
	-rm -f ./$(releaseName)

