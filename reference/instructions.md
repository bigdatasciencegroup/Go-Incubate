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