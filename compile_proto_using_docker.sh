#!/bin/bash

ALAMEA_GRPC_GO_IMAGE_REPO="alameda/grpc_go_$(docker run --rm -v $(pwd):$(pwd) -w $(pwd) golang:stretch git rev-parse --abbrev-ref HEAD)"
ALAMEA_GRPC_GO_IMAGE_TAG="latest"
ALAMEA_GRPC_GO_IMAGE="$ALAMEA_GRPC_GO_IMAGE_REPO:$ALAMEA_GRPC_GO_IMAGE_TAG"
ALAMEA_GRPC_GO_IMAGE_DOCKERFILE=Dockerfile_gRPC_go

generate_dockerfiles(){
    cat > $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE - <<EOF
FROM golang:stretch
ARG DOCKERFILE_MD5
ENV DOCKERFILE_MD5=\$DOCKERFILE_MD5 PROTOC_VER=4.0.0 OS_ARC=linux-x86_64 PROTOC_GEN_GO_VER=v1.4.2 PROTOC_GEN_DOC_VER=v1.3.2
COPY setup.py .
RUN apt-get update && apt-get install unzip python3 python3-pip -y && \\
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v\$PROTOC_VER-rc2/protoc-\$PROTOC_VER-rc-2-\$OS_ARC.zip && \\
unzip protoc-\$PROTOC_VER-rc-2-\$OS_ARC.zip -d /usr/local && rm protoc-\$PROTOC_VER-rc-2-\$OS_ARC.zip && \\
go get -d -u github.com/golang/protobuf/protoc-gen-go && \\
git -C "\$(go env GOPATH)"/src/github.com/golang/protobuf checkout \$PROTOC_GEN_GO_VER && \\
go install github.com/golang/protobuf/protoc-gen-go && \\
GO111MODULE=on go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@\$PROTOC_GEN_DOC_VER && \\
perl -lne 'print if /INSTALL_REQUIRES =/ .. /\)/' setup.py | grep -v \( | grep -v \) |  awk -F  "'" '{print \$2}' > requirements.txt && \\
pip3 install -r requirements.txt && rm requirements.txt && rm setup.py && \\
rm -rf /var/lib/apt/lists/*
EOF
}

build_go_image(){
    local DOCKERFILE_MD5=`md5sum $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE | awk '{print $1}'`
    docker build --build-arg DOCKERFILE_MD5=$DOCKERFILE_MD5 . -t $ALAMEA_GRPC_GO_IMAGE -f $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE 
}

compile_grpc(){
    echo "Check gRPC go image."
    local DOCKERFILE_MD5=`md5sum $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE | awk '{print $1}'`
    if ! docker images $ALAMEA_GRPC_GO_IMAGE | grep $ALAMEA_GRPC_GO_IMAGE_REPO > /dev/null 2>&1; then
        echo "Build new image $ALAMEA_GRPC_GO_IMAGE";
        build_go_image
    fi
    if ! docker run --rm $ALAMEA_GRPC_GO_IMAGE sh -c "
        if [ \"$DOCKERFILE_MD5\" != \"\$DOCKERFILE_MD5\" ]; then
            exit 1;
        fi"; then
        echo "Refresh image $ALAMEA_GRPC_GO_IMAGE";
        docker rmi $ALAMEA_GRPC_GO_IMAGE;
        build_go_image
    fi
    echo "Start compiling proto files."
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) $ALAMEA_GRPC_GO_IMAGE bash -c "for pt in \$(find . | grep \\\.proto\$ | grep -v ^\\\./include | grep -v ^\\\./google);do protoc -I . -I include/ \$pt --go_out=paths=source_relative,plugins=grpc:. --js_out=import_style=commonjs,binary:.; python3 -m grpc_tools.protoc -I . -I include/ --python_out=./ --grpc_python_out=./ \$pt; done"
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) $ALAMEA_GRPC_GO_IMAGE bash -c "protoc -I . -I include/ --doc_out=./ --doc_opt=html,index.html $(find . | grep \\\.proto\$ | grep -v ^\\\./include | grep -v ^\\\./google | tr '\n' ' ');"
    echo "Finish compiling proto files."
}

remove_dockerfiles(){
   [ -f $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE ] && rm $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE
}

generate_dockerfiles
compile_grpc
remove_dockerfiles
