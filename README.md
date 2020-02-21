# Linstor Operator

This is the initial public alpha for the Linstor Operator. Currently, it is
suitable for testing and development.

## Contributing

This Operator is currently under heavy development: documentation and examples will chage frequently. 
Always use the latest release.

If you'd like to contribute, please visit https://gitlab.com/linbit/linstor-operator
and look through the issues to see if there something you'd like to work on. If
you'd like to contribute something not in an existing issue, please open a new
issue beforehand.

If you'd like to report an issue, please use the issues interface in this
project's gitlab page.

## Building and Development

This project is built using the operator-skd (version 0.13.0). Please refer to
the [documentation for the sdk](https://github.com/operator-framework/operator-sdk/tree/v0.9.x)

## Deployment with Helm v3 Chart

The operator can be deployed with Helm v3 chart in /charts.
- Label the worker nodes with `linstor.linbit.com/linstor-node=true` such as;
  ```
  kubectl label no my-worker-node linstor.linbit.com/linstor-node=true
  ```

- Create a kubernetes secret to allow obtaining LINSTOR images from the 
  drbd.io repo.  This example will create a secret named `drbdiocred`:
  ```
  kubectl create secret docker-registry drbdiocred --docker-server=drbd.io --docker-username=<YOUR LOGIN> --docker-email=<YOUR EMAIL> --docker-password=<YOUR PASSWORD>
  ```
  Then this secret name must be specified in the charts/linstor/values.yaml.  
  ```
  drbdRepoCred: drbdiocred  # <- Specify the kubernetes secret name
  ```

- Configure the LVM VG and LV names in charts/linstor/values.yaml.
  ```
  lvmPoolVgName: "drbdpool"      # <- Local VG name for Thick LVM Pool
  lvmThinPoolVgName: "drbdpool"  # <- Local VG name for Thin LVM Pool
  lvmThinPoolLvName: "thinpool"  # <- Local LV name for Thin LVM Pool
  ```

- Finally create a Helm release named `linstor-op` that will set up
  everything.
  ```
  helm install linstor-op ./charts/linstor -f ./charts/linstor/values.yaml
  ```
### Terminating Helm release/deployment

```
helm delete linstor-op
```
will terminate the Linstor release.  However due to the Helm's current policy,
the newly created Custom Resource Definitions named linstorcontrollerset and
linstornodeset will __not__ be deleted by the command.  This will also cause
the instances of those CRD's named `linstor-op-linstor-ns` and `linstor-op-linstor-cs`
to remain running.

To terminate those instances after the helm delete command, run
```
kubectl patch linstorcontrollerset linstor-op-linstor-cs -p '{"metadata":{"finalizers":[]}}' --type=merge

kubectl patch linstornodeset linstor-op-linstor-ns -p '{"metadata":{"finalizers":[]}}' --type=merge
```

After that, all the instances created by the Helm release will be terminated.

More information regarding Helm's current position on CRD's can be found
[here](https://helm.sh/docs/topics/chart_best_practices/custom_resource_definitions/#method-1-let-helm-do-it-for-you).

## Deployment without using Helm v3 chart

### Configuration

The operator must be deployed within the cluster in order for it it to have access
to the controller endpoint, which is a kubernetes service. See the operator-sdk
guide on the Operator Framework for more information.

Worker nodes will only run on kubelets labeled with `linstor.linbit.com/linstor-node=true`
```
kubectl label no my-worker-node linstor.linbit.com/linstor-node=true
```

### Etcd

An etcd cluster must be running and reachable to use this operator. By default,
the controller will try to connect to `linstor-etcd` on port `2379`

A simple in-memory etcd cluster can be set up using helm:
```
kubectl create -f examples/etcd-env-vars.yaml
helm repo add bitnami https://charts.bitnami.com/bitnami
helm install linstor-etcd bitnami/etcd --set statefulset.replicaCount=3 -f examples/etcd-values.yaml
```

If you are using Helm 2 and encountering difficulties with the above steps, you may need to set RBAC
rules for the tiller component of helm:
```
kubectl create serviceaccount --namespace kube-system tiller
kubectl create clusterrolebinding tiller-cluster-rule --clusterrole=cluster-admin --serviceaccount=kube-system:tiller
kubectl patch deploy --namespace kube-system tiller-deploy -p '{"spec":{"template":{"spec":{"serviceAccount":"tiller"}}}}'
```

### Kubernetes Secret for drbd.io Repo Access

Create a kubernetes secret to allow obtaining LINSTOR images from the
drbd.io repo.  Create a secret named `drbdiocred` like this;
```
kubectl create secret docker-registry drbdiocred --docker-server=drbd.io --docker-username=<YOUR LOGIN> --docker-email=<YOUR EMAIL> --docker-password=<YOUR PASSWORD>
```

### Deploy Operator

Inspect the basic deployment example (examples/linstor-operator-part-1.yaml), it may be deployed by:
```
kubectl create -f examples/linstor-operator-part-1.yaml
```
Lastly, edit the storage nodes' LVM VG and LV names in examples/linstor-operator-part-2.yaml.  Now you can finally deploy the LINSTOR cluster with:

```
kubectl create -f examples/linstor-operator-part-2.yaml
```

## License

Apache 2.0
