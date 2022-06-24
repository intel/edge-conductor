# Edge Conductor

## Welcome to Edge Conductor

Edge Conductor is a framework to create, manage and operate Kubernetes clusters
for Internet of Things (IOT) workloads. Essentially, Edge Conductor is a
software reference stack for IOT Edge Computing. It provides orchestration
capabilities for the infrastructure edge and for the device edge, while enabling
common management for infrastructure systems and also a wide range of devices
and endpoints.

## Contents of this guide


  * [Edge Conductor Documentation](#edge-conductor-documentation)
  * [Edge Conductor Deployment Models](#edge-conductor-deployment-models)
  * [Download and Build Edge Conductor Code (Internal users only)](#download-and-build-edge-conductor-code-internal-users-only)
  * [HW Requirements for Edge Conductor Day-0 Host](#hw-requirements-for-edge-conductor-day-0-host)
  * [OS and System Requirements for Edge Conductor Day-0 Host](#os-and-system-requirements-for-edge-conductor-day-0-host)
  * [FAQ](#faq)



## Edge Conductor Documentation

The following links will guide you through the Edge Conductor documents,
including running a simple Get Started application and trying some tutorials.

*  First time user? Use the [Get Started](docs/guides/get-started.md) guide to
   use the KIND cluster and default configuration files.

*  Use the Edge Conductor [Tutorials](docs/tutorials/index.md) to learn what Edge
   Conductor is and what you can do with it. You will learn how to deploy some
   simple applications on the Kubernetes cluster you built with Edge Conductor.

   * [Example: Hello Cluster!](docs/tutorials/samples/hello-cluster.md)
   * [Example: Hello Cluster! Helm Version](docs/tutorials/samples/hello-cluster-helm.md)
   * [Example: Web Indexing Sample Application](docs/tutorials/samples/web-indexing.md)


*  Learn about sample [Edge Conductor
   Configurations](docs/guides/ec-configurations.md), which are provided in
   `_workspace/experienceKit`. You can modify configurations and use specific
   config files to run the Edge Conductor tool.

*  Deploy a Target Cluster using the Edge Conductor tool. See the following
   guides for detailed instructions for different cluster types:

   * [Deploy a KIND Cluster](docs/guides/cluster-deploy-KIND.md)
   * [Deploy an RKE Cluster](docs/guides/cluster-deploy-RKE.md)
   * [Deploy a Cluster with ClusterAPI](docs/guides/cluster-deploy-ClusterAPI.md)



## Edge Conductor Deployment Models

Edge Conductor can be used to deploy Kubernetes clusters in several different
ways. All of the deployment methods rely on a "Day-0" host machine, which is a
dedicated system that runs the Edge Conductor tools that download, build and
provision the required software. (*Day-0* is a lifecycle term used in network
automation and cloud deployment, where the high-level task on Day-0 is
installation, Day-1 is configuration, and Day-2 is optimization. Here it
indicates the host system used for installing the Edge Conductor tools.)


All Edge Conductor deployment methods also rely
on a management console to operate and administer the cluster.

  * Virtual cluster: All nodes in the cluster are virtualized as container
    images and Kubernetes is deployed on the user’s local machine. The Day-0
    host and management console are also the user's local machine. This
    deployment model is implemented using "kind" as described below.

  * On-premise cluster: All nodes in the cluster are actual physical (or
    virtual) machines.  The Day-0  and management console functions are
    performed on machine(s) outside the cluster.  This deployment model is
    implemented using either Rancher RKE or the Cluster API as described below.

  * Existing cluster: All nodes in the cluster have already been provisioned
    and are running Kubernetes, for example as in a StarlingX, VMWare Tanzu or
    OpenShift cluster or similar Kubernetes deployment. In this model only
    services are deployed and managed by the Edge Conductor tools. The Day-0 machine
    and management cluster are assumed to already exist.




## Download and Build Edge Conductor Code

1.  Get the code using one of the following methods:

    * Get the code from git repo:
        
        ```bash
        git clone https://github.com/intel/edge-conductor.git edge-conductor
        ```

    * Unzip the code package, for example, `edge_conductor_<version>.tar.gz`:

        ```bash
        tar zxvf edge_conductor_<version>.tar.gz
        ln -sf edge_conductor_<version> edge-conductor
        ```

2.  Build the code:

    ```
    cd edge-conductor
    make
    ```

    You will see output similar to:

    ```
    make -C api/schemas
    make -C api/proto build
    go mod tidy
    go run build/plugingenerator.go
    go vet ./pkg/... ./cmd/...
    Going to format and build code.
    ...
    go build -v -o _workspace/bin/conductor cmd/conductor/main.go
    ...
    ```

    If `make` is successful, no ERROR messages are displayed on the console.

    With the `make` command, binary files are generated under `_workspace`
    folder and configuration files are also copied to `_workspace`.

3.  Change to the `_workspace` directory, which is created after `make`
    completes:

    ```
    cd _workspace
    ```

    The file structure generated under the `_workspace` folder is:

    ```
    _workspace/
    ├── bin
    ├── config
    ├── experienceKit
    ├── conductor -> bin/conductor
    ├── kubectl -> bin/kubectl
    └── workflow
    ```

Run ``./conductor help`` to see command line usage.

Refer to [Get Started](docs/guides/get-started.md) to continue with common tasks
such as deploying a target cluster, deploying services on the target cluster,
and stopping Edge Conductor services.




## HW Requirements for Edge Conductor Day-0 Host

The Day-0 host should meet the following minimal hardware requirements:

- CPU: 2 or more cores
- Memory: 2 gigabytes (GB) or greater.
- Storage: 10 GB or greater available storage is required to build and run the
  Day-0 environment.
- Internet connection: Internet connectivity is necessary to download and use
  some features.


Node Requirements:

- CPU: 2 or more cores
- Memory: 2 gigabytes (GB) or greater.
- Storage: 10 GB or greater available storage is required to build and run the
  Day-0 environment.
- Internet connection: Internet connectivity is necessary to download and use
  some features.



## OS and System Requirements for Edge Conductor Day-0 Host

- Ubuntu 18.04+
- make 4.1+
- unzip 2.11+
- DockerCE
    * 20.10.3+ (for DockerCE v20)
    * After you install DockerCE, configure a user group so you can use Docker
      without the `sudo` command.
      For details, follow [Post-installation steps for Linux](https://docs.docker.com/engine/install/linux-postinstall/).
- git 2.33.0+
- sudo
    * Passwordless sudo should be configured on every worker node. For details, follow [How to setup passwordless sudo](https://serverfault.com/questions/160581/how-to-setup-passwordless-sudo-on-linux)


Additional software:

*   Install [kubectl (v1.20.0)](https://kubernetes.io/docs/tasks/tools/) to
    interact with the clusters created with Edge Conductor. For details, follow
    the Kubernetes steps: [Install and Set Up kubectl on Linux](https://kubernetes.io/docs/tasks/tools/install-kubectl-linux/).


Proxy setup:

*   Be sure to set your proxies correctly. Some environments require network
    proxies for Docker operations (e.g. docker pull, docker push, docker run,
    and so on).
*   For http_proxy/HTTP_PROXY and https_proxy/HTTPS_PROXY, you must test they
    are actually working. Also note that the no_proxy/NO_PROXY list **MUST NOT
    contain spaces** between the addresses.



## FAQ

Frequently asked questions.

### How do I access the Portainer Web UI on KIND cluster?

The latest KIND release doesn't support LoadBalancer, so the External-IP for the
Portainer service Web UI will always be pending as shown below:

```
$ kubectl get pod,svc -n portainer
NAME                            READY   STATUS    RESTARTS   AGE
pod/portainer-fc654c454-g4qvl   1/1     Running   26         3h26m

NAME                TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                         AGE
service/portainer   LoadBalancer   10.96.182.150   <pending>     9000:31343/TCP,8000:32344/TCP   3h26m
```

To access the Portainer service in this situation, use the `kubectl
port-forward` command to forward a local port to Portainer service port 9000,
which is used for the Web UI.

For example, the following command will forward the access request from
localhost 5990 port to Portainer service at 9000:

```
$ kubectl port-forward -n portainer service/portainer 5990:9000
Forwarding from 127.0.0.1:5990 -> 9000
```

When the port-forwarding is running (shown by the log message `Forwarding
from ...`), you can access the 127.0.0.1:5990 to visit the Portainer web UI and
create the initial admin account.

***NOTE***: If there's no client accessing the forwarded port for a while (50 seconds), the port-forwarding will stop automatically.

***NOTE***: If you do not create your admin account in 5 minutes after the service is launched, then the Portainer service will restart itself automatically. (See the Portainer issue reported here: https://github.com/portainer/portainer/issues/2475)



Copyright (C) 2022 Intel Corporation

SPDX-License-Identifier: Apache-2.0
