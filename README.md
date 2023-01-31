# emkr
emkr is **E**xtract the **M**anifest for a specific **K**ubernetes **R**esource


## Usage

Extract Kubernetes manifests that match the specified key and value.


```
$ cat multiple_documents.yaml | emkr -k kind -v Deployment

> ---
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: hoge-deployment
    labels:
      app: hoge
  spec:
    replicas: 3
    selector:
      matchLabels:
        app: hoge
    template:
      metadata:
        labels:
          app: hoge
      spec:
        containers:
        - name: apps
          image: hogeapp:1.0
  ---
  apiVersion: apps/v1
  kind: Deployment
  metadata:
    name: fuga-deployment
    labels:
      app: fuga
  spec:
    replicas: 3
    selector:
      matchLabels:
        app: fuga
    template:
      metadata:
        labels:
          app: fuga
      spec:
        containers:
        - name: apps
          image: fugaapp:1.0
```
```
$ cat multiple_documents.yaml | emkr -k metadata.labels.app -v hoge

>  ---
   apiVersion: apps/v1
   kind: Deployment
   metadata:
     name: hoge-deployment
     labels:
       app: hoge
   spec:
     replicas: 3
     selector:
       matchLabels:
         app: hoge
     template:
       metadata:
         labels:
           app: hoge
       spec:
         containers:
         - name: apps
           image: hogeapp:1.0
```
