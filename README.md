# monitoring
Metrics ingestion server.

## Overview

This is a metrics ingestion server written in Golang for ingesting metrics of clients and generating an on demand report for maximum CPU and memory usage.

The server exposes the following REST API's -

- `/metrics` - This API is used for ingesting metrics of clients. A sample curl request for ingesting metrics looks like below (assuming our api is deployed locally and exposed to port 8080) -

```
curl \
    -XPOST \
    -H "Content-Type: application/json" \
    --data '{"percentage_cpu_used": 55, "percentage_memory_used": 90}' \
    http://127.0.0.1:8080/metrics
```
- `/report` - This API generates an on demand report with the max CPU and memory usage that the client ever reached. A sample curl request for generating the report looks like below -

```
curl \
    -XGET \
    -H "Content-Type: application/json" \
    http://127.0.0.1:8080/report
```

## Building a Docker image

- Clone the repository -

  ```
  git clone https://github.com/harshshekhar15/monitoring.git
  ```

- Get inside the project -

  ```
  cd monitoring
  ```

- Run build command -

  ```
  docker build -t <ORG-NAME>/<REPO-NAME>:<IMAGE-TAG> .
  ```

## Spinning up the server

To quickly spin up the metrics server in a docker container, run the below command -

```
docker run -d -p 8080:8080 --name <CONTAINER-NAME> <ORG-NAME>/<REPO-NAME>:<IMAGE-TAG>
```

After running the above command metrics server will be running at `http://127.0.0.1:8080`