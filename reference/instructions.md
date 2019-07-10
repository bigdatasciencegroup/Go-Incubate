# Instructions

1. To generate `requirements.txt` file for Python Docker containers
    ```python
        pipreqs [options] <path to main project folder>

        [options]
        --force : to overwrite existing file
    ```
1. To install dependencies
    ```python
        pip install -r /app/requirements.txt
    ```

1. To visualize Kafka inside docker, issue the following command:

   ```text
   docker run -e ADV_HOST=127.0.0.1 -e EULA="https://dl.lenses.stream/d/?id=8a668761-4d2b-4f23-bb84-0c5f9964d772" --rm -p 3030:3030 -p 9092:9092 -p 2181:2181 -p 8081:8081 -p 9581:9581 -p 9582:9582 -p 9584:9584 -p 9585:9585 landoop/kafka-lenses-dev
   ```

1. To build website on localhost
    ```text
    bundle exec jekyll serve --watch --incremental
    ```

1. Transfer docker image without using internet based repository
    You will need to save the Docker image as a tar file:

    docker save -o <path for generated tar file> <image name>
    Then copy your image to a new system with regular file transfer tools such as cp or scp. After that you will have to load the image into Docker:

    docker load -i <path to image tar file>
    PS: You may need to sudo all commands.

    EDIT: You should add filename (not just directory) with -o, for example:

    + docker save -o c:/myfile.tar centos:16
    + docker save -o "C:\Projects\goWorkspace\src\github.com\adaickalavan\dockerimages\zookeeper.tar" confluentinc/cp-zookeeper

1. To see git hitsory, run:
    ```git
    git log --graph --oneline --decorate
    ```

1. To profile a Go code using a web interface, run:
    ```go
    go tool pprof -http=localhost:8080 /tmp/cpu.pprof
    go tool pprof -http=localhost:8080 /tmp/mem.pprof
    ```

1. To see time spent on each line, run:
    ```go
    go tool pprof /tmp/cpu.pprof
    list <packageName>
    ```

1. Show tensorflow saved model parameters
    ```go
    saved_model_cli show --dir ~/GoWorkspace/src/github.com/adaickalavan/Scalable-Deployment-Kubernetes/tfserving/resnet/1538687457 --all
    ```

1. Start bash inside a docker container
    ```go
    ~/GoWorkspace/src/github.com/adaickalavan/Scalable-Deployment-Kuberserving/resnet/1538687457$ docker exec -it goconsumer bash   
    ``` 
    alternative command
    ```
    docker run -it cvideoingest /bin/bash
    ```
1. To visualize Godoc of our own package    