.PHONY: run build image push clean

tag = v1.5
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

run-web:

	docker rm -f web
	docker rmi  ${web-image}:$(tag)
	docker run -d --name web -p 80:80 -v /Users/lei/Documents/GitHub/lbemi/nginx:/etc/nginx/conf.d ${web-image}:$(tag)
