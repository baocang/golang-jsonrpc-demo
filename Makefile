main:
	docker build -t baocang/golang-jsonrpc-demo:0.0.1 .
up:
	kubectl apply -f ./k8s/namespace.yml
	kubectl apply -f ./k8s/service.yaml
	kubectl apply -f ./k8s/gateway.yml
down:
	kubectl delete -f ./k8s/gateway.yml
	kubectl delete -f ./k8s/service.yaml
	kubectl delete -f ./k8s/namespace.yml
debug:
	telepresence intercept --mount true --port 8080 --env-file .cluster.env --namespace develop golang-jsonrpc
leave:
	telepresence leave golang-jsonrpc-develop
