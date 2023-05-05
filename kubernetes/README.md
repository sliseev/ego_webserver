Start k8s cluster
=================

```
$ sudo docker build -t egoserv .
$ sudo docker tag egoserv <user>/egoserv

$ # Upload image to https://hub.docker.com/
$ sudo docker login -u <user>
$ sudo docker push <user>/egoserv

$ cd kubernetes
$ minikube start --driver=kvm2 (on linux: https://minikube.sigs.k8s.io/docs/drivers/kvm2/)
  or minikube start --driver=hyperkit (on MacOS: brew install hyperkit)

$ kubectl apply -f postgres-secret.yaml
$ kubectl apply -f postgres-configmap.yaml
$ kubectl apply -f postgres.yaml
$ kubectl apply -f ego-server.yml
$ kubectl get pod
$ kubectl get svc

$ minikube service ego-server-service
$ curl http://<external-ip>:30080/driver
 ```

If pod doesn't start use the following commands for troubleshooting:
```
$ kubectl describe pod <podname>
$ kubectl logs <podname>
$ kubectl exec -it <podname> -- /bin/bash
```

Playing with cluster:
```
$ kubectl scale --replicas=2 deployment/ego-server
```
