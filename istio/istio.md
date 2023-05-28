# 安装


### 虚拟服务路径重写

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: VirtualService
metadata:
  name: prodvs
  namespace: myistio
spec:
  hosts:
    - prodsvc
  http:
  - match:
    - uri:
       prefix: "/p" # 前缀是p路径重写为/prods
    rewrite:
       uri: "/prods"
    route:
     - destination:
        host: prodsvc  # 目标service
        subset: v1svc # 指定destination rule为v1svc版本
        port:
          number: 80
  - match:
    - uri:
       prefix: "/r"
    rewrite: 
       uri: "/reviews"
    route:
     - destination:
        host: reviewsvc
        port:
          number: 80
```
测试
```bash
# 请求地址会被重写 prodsvc/prods/123
/ # curl prodsvc/p/123
{"Id":123,"Name":"测试商品","Reviews":[{"Id":2000,"Title":"测试评论1"},{"Id":2002,"Title":"测试评论2"}]}/ # 
```

### Destination Rule

虚拟服务将流量如何路由到给定目标 做了定义
而 Destination Rule用于配置将流量转发到实际工作负载时应用的策略（譬如流量拆分、灰度发布、负载均衡 等 ）,需要和vs配合使用

```yaml
apiVersion: networking.istio.io/v1alpha3
kind: DestinationRule
metadata:
  name: prod-rule
  namespace: myistio
spec:
  host: prodsvc
  subsets:
    - name: v1svc
      labels:
        version: v1
    - name: v2svc
      labels:
        version: v2
```
#### TrafficPolicy
