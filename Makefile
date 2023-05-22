.PHONY: run build image push clean

tag = v0.1
server-image = registry.cn-hangzhou.aliyuncs.com/lbemi/crmdb-server
web-image = registry.cn-hangzhou.aliyuncs.com/lbemi/crmdb-web

ALL: run

run: build
	./crmdb

build-server:
	cd ./server && docker build -t ${server-image}:$(tag) .
push-server: 
	docker push  ${server-image}:$(tag)

build-web:
	cd ./vue-next-admin && docker build -t ${web-image}:$(tag) .
push-web: 
	docker push  ${web-image}:$(tag)


