## operator_sdk_pro1

```
operator-sdk new operator-sdk-pro1
operator-sdk add api --api-version=branch.dzht.com/v1 --kind=AdminInfo
operator-sdk add controller --api-version=branch.dzht.com/v1 --kind=AdminInfo
```

Update `pkg/apis/branch/v1/admininfo_types.go` to add fields.

```
operator-sdk generate k8s
```

Update `pkg/controller/admininfo/admininfo_controller.go` to add logic.

```
operator-sdk build wanghaihui/admin-operator:v0.0.1
```

Update `deploy/operator.yaml` to replace <REPLACE_IMAGE>.

```
docker push wanghaihui/admin-operator:v0.0.1
```

```
kubectl apply -f deploy/crds/branch.dzht.com_admininfos_crd.yaml
kubectl apply -f deploy/role.yaml
kubectl apply -f deploy/service_account.yaml
kubectl apply -f deploy/role_binding.yaml
kubectl apply -f deploy/operator.yaml
```

```
kubectl delete -f deploy/operator.yaml
kubectl delete -f deploy/role_binding.yaml
kubectl delete -f deploy/service_account.yaml
kubectl delete -f deploy/role.yaml
kubectl delete -f deploy/crds/branch.dzht.com_admininfos_crd.yaml
```

```
kubectl get AdminInfo
kubectl apply -f deploy/crds/branch.dzht.com_v1_admininfo_cr.yaml
kubectl delete -f deploy/crds/branch.dzht.com_v1_admininfo_cr.yaml
```