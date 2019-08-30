---
title: "Docker"
permalink: /docs/docker/
toc: true
---

## Docker File
+ Containerization is packaging an application, its dependencies, and its configuration into a single deployable unit, called an image.
+ Failed Docker containers can be restarted automatically.
+ Can be run on Linux, Mac, Windows
+ Aids reusability, refactorability, scalability, and maintainability of code.
+ Each command used to generate Docker image  generates a new intermediate image via a delta — essentially capturing only what changed from the previous intermediate step.
+ The following is a representative Docker file with multistage build feature. Build the image using the command `$ docker build -t <container name>` .
<br/><img src="/assets/images/wiki/wiki_docker_01.jpg" width="100%"/>

## Docker Compose
+ To instantiate the container image, run a docker compose file by executing `$ docker-compose up`. Docker compose file contains instructions to instantiate the image, set environment variables, setup networks, etc. 
<br/><img src="/assets/images/wiki/wiki_docker_02.jpg" width="100%" height="160%"/>

## Configuration
+ Need to set proxy for Docker if we use Docker behind a corporate proxy, e.g., `http://10.0.0.0:8080/`. Hence, create a `http-proxy.conf` file at `/etc/systemd/system/docker.service.d/` such that `/etc/systemd/system/docker.service.d/http-proxy.conf` file contains the following.
    ```
    [Service]
    Environment="HTTP_PROXY=http://10.0.0.0:8080/"
    Environment="HTTPS_PROXY=https://10.0.0.0:8080/"
    Environment="NO_PROXY=http://172.16.0.0/"
    ```    
    Note that we can set no proxy for internal IP addresses such as `http://172.16.0.0/`.

+ Restart the Docker daemon to start using the new proxy setting.
    ```bash
    $ systemctl daemon-reload
    $ service docker restart
    ```