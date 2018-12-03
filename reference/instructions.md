# Instructions

1. To generate `requirements.txt` file for Python Docker containers
    ```python
        pipreqs [options] <path to main project folder>

        [options]
        --force : to overwrite existing file
    ```
1. To visualize Kafka inside docker, issue the following command:

   ```text
   docker run -e ADV_HOST=127.0.0.1 -e EULA="https://dl.lenses.stream/d/?id=8a668761-4d2b-4f23-bb84-0c5f9964d772" --rm -p 3030:3030 -p 9092:9092 -p 2181:2181 -p 8081:8081 -p 9581:9581 -p 9582:9582 -p 9584:9584 -p 9585:9585 landoop/kafka-lenses-dev
   ```

1. To build website on localhost
    ```text
    bundle exec jekyll serve --watch --incremental
    ```

1. Demo:
    1. Restart the video from where it ended previously after a video crash
    1. Play two video streaming from the same queue using two different consumer group

1. Transfer docker image without using internet based repository
    You will need to save the Docker image as a tar file:

    docker save -o <path for generated tar file> <image name>
    Then copy your image to a new system with regular file transfer tools such as cp or scp. After that you will have to load the image into Docker:

    docker load -i <path to image tar file>
    PS: You may need to sudo all commands.

    EDIT: You should add filename (not just directory) with -o, for example:

    docker save -o c:/myfile.tar centos:16


1. 