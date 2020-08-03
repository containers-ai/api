# api

This repository defines interfaces for Alameda to access data sources and store its generated predictions and recommendations. The backend that implements these interfaces is the *datahub* component of Alameda. See [Alameda architecture design](https://github.com/containers-ai/alameda/blob/master/design/architecture.md) for more details.

## How to compile

We provide two ways to compile proto files: [within docker environment](#compile-within-docker-environment) and [without docker environment](#compile-without-docker-environment).

### Compile within docker environment

Run the following script to compile proto files with docker
```bash
./compile_proto_using_docker.sh
```
The generated code will be located in the same folder as the .proto files.

## How to use

### Coding with golang

Add the following import declarations in your .go files when using the Alameda API gRPC calls.
```
import "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
```

## Coding with python

Install alameda-api packages by
```bash
pip install git+https://github.com/containers-ai/api.git
```
Then you can use Alameda API gRPC calls in your .py files by
```
from alameda_api.v1alpha1.datahub import server_pb2, server_pb2_grpc
```
