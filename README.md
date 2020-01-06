## operator_sdk_pro1
```
operator-sdk new operator_sdk_pro1
operator-sdk add api --api-version=branch.dzht.com/v1 --kind=AdminInfo
operator-sdk add controller --api-version=branch.dzht.com/v1 --kind=AdminInfo
```

Update `pkg/apis/branch/v1/admininfo_types.go` to add fields.

```
operator-sdk generate k8s
```

Update `pkg/controller/admininfo/admininfo_controller.go` to add logic.

```
operator-sdk build wanghh2000/admin-operator:v0.0.1
```

Update `deploy/operator.yaml` to replace <REPLACE_IMAGE>.