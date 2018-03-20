kubectl create ns entrypoint || true
kubectl create -f kubernetes/frontier-config.yaml --namespace=entrypoint
kubectl create -f kubernetes/service.yaml --namespace=entrypoint
kubectl create -f kubernetes/deployment.yaml --namespace=entrypoint
