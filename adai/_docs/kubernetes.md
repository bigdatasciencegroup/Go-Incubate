---
title: "Kubernetes"
permalink: /docs/kubernetes/
---

+ Container orchestration system: automates container deployment, scaling, networking, storage, and scheduling.
+ `node`: A worker machine in the Kubernetes cluster, responsible for actually running pods
+ `pod`: layer of abstraction containing one or more container runtimes like Docker or rkt. Single-container pods are preferred.
+ If a machine within the cluster dies, affected container will be migrated to another machine. 
+ `kubectl`: A command-line utility with commands to control the Kubernetes cluster

+ Commands:
    ```bash
    $ docker build -t <container name> . #To build the container
    $ kubectl <command> <resource> <options> #General command format
    $ kubectl get pods <name of resource> #Retrieve information about a pod
    $ kubectl get all #Summary of each resource in cluster
    $ kubectl describe all #Provides detailed information
    $ kubectl apply -f <file name>.yml #To create, delete, or replace state of Kubernetes cluster
    $ kubectl delete -f <file name>.yml #To delete the object declared by the YAML file 
    ```
+ Kubernetes object Kind: Pod, ConfigMap, Deployment, Service

## Kubernetes Deployment
+ Deployment manages, replicates to handle increased load, and updates, a set of pods through a single parent object
+ Nests a pod specification (deployment.spec.template) inside the deployment's configuration 
+ A representative Kubernetes Deployment file is shown below.
<br/><img src="/assets/images/wiki/wiki_kubernetes_01.jpg" width="100%"/>

+ `Horizontal Pod Autoscaler` is a new resource in modern versions of Kubernetes that manages the number of replicas in your deployment automatically, based on resource utilization(e.g., memory, CPU, custom metrics).
<br/><img src="/assets/images/wiki/wiki_kubernetes_03.jpg" width="30%"/>

## Kubernetes Service
+ Services routes network requests to appropriate pods based on matching labels
+ A representative `Kubernetes Service` file is shown below.
<br/><img src="/assets/images/wiki/wiki_kubernetes_02.jpg" width="100%"/>

+ The `Kubernetes Service` in the above file is illustrated pictorially below.
<br/><img src="/assets/images/wiki/wiki_kubernetes_04.jpg" width="100%" height="150%"/>
