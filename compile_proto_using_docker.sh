#!/bin/bash

ALAMEA_GRPC_GO_IMAGE_REPO="alameda/grpc_go_$(docker run --rm -v $(pwd):$(pwd) -w $(pwd) golang:stretch git rev-parse --abbrev-ref HEAD)"
ALAMEA_GRPC_GO_IMAGE_TAG="latest"
ALAMEA_GRPC_GO_IMAGE="$ALAMEA_GRPC_GO_IMAGE_REPO:$ALAMEA_GRPC_GO_IMAGE_TAG"
ALAMEA_GRPC_GO_IMAGE_DOCKERFILE=Dockerfile_gRPC_go

generate_dockerfiles(){
    cat > $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE - <<EOF
FROM golang:stretch
ARG DOCKERFILE_MD5
ENV DOCKERFILE_MD5=\$DOCKERFILE_MD5 PATH="/usr/local/node/bin:/go/bin:${PATH}"
ENV PROTOC_VER=4.0.0 OS_ARC=linux-x86_64 PROTOC_GEN_GO_VER=v1.4.2 PROTOC_GEN_DOC_VER=v1.3.2 API_COMMON_PROTOS_VER=1.50.0 PROTOC_GEN_WEB_VER=1.2.1 NODE_VER=v14.8.0
COPY setup.py .
RUN apt-get update && apt-get install unzip xz-utils python3 python3-pip -y && \\
curl -LO https://github.com/protocolbuffers/protobuf/releases/download/v\$PROTOC_VER-rc2/protoc-\$PROTOC_VER-rc-2-\$OS_ARC.zip && \\
unzip protoc-\$PROTOC_VER-rc-2-\$OS_ARC.zip -d /usr/local && rm protoc-\$PROTOC_VER-rc-2-\$OS_ARC.zip && \\
#curl -LO https://github.com/grpc/grpc-web/releases/download/\$PROTOC_GEN_WEB_VER/protoc-gen-grpc-web-\$PROTOC_GEN_WEB_VER-linux-x86_64 && \\
#mv protoc-gen-grpc-web-\$PROTOC_GEN_WEB_VER-linux-x86_64 /usr/local/bin/protoc-gen-grpc-web && chmod +x /usr/local/bin/protoc-gen-grpc-web && \\
curl -LO https://nodejs.org/dist/\$NODE_VER/node-\$NODE_VER-linux-x64.tar.xz && tar Jxvf node-\$NODE_VER-linux-x64.tar.xz && \\
mv node-\$NODE_VER-linux-x64 /usr/local/node && rm node-\$NODE_VER-linux-x64.tar.xz && \\
go get -d -u github.com/golang/protobuf/protoc-gen-go && \\
git -C "\$(go env GOPATH)"/src/github.com/golang/protobuf checkout \$PROTOC_GEN_GO_VER && \\
go install github.com/golang/protobuf/protoc-gen-go && \\
GO111MODULE=on go get -u github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@\$PROTOC_GEN_DOC_VER && \\
git clone --depth 1 --branch \$API_COMMON_PROTOS_VER https://github.com/googleapis/api-common-protos.git && \\
mv api-common-protos/google/rpc /usr/local/include/google/ && rm -rf api-common-protos && \\
perl -lne 'print if /INSTALL_REQUIRES =/ .. /\)/' setup.py | grep -v \( | grep -v \) |  awk -F  "'" '{print \$2}' > requirements.txt && \\
pip3 install -r requirements.txt && rm requirements.txt && rm setup.py && npm i -g --unsafe-perm grpc-tools protoc-gen-grpc && \\
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

    echo "Remove generated files."
    rm -f `find .| grep -E '\.go$|\.py$|\.js$|\.html$' | grep -v setup.py`
    echo "Start compiling proto files."
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) $ALAMEA_GRPC_GO_IMAGE bash -c "for pt in \$(find . | grep \\\.proto\$);do grpc_tools_node_protoc -I . -I /usr/local/include --js_out=import_style=commonjs,binary:. --grpc_out=. --plugin=protoc-gen-grpc=/usr/local/node/lib/node_modules/grpc-tools/bin/grpc_node_plugin \$pt; done"
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) $ALAMEA_GRPC_GO_IMAGE bash -c "for pt in \$(find . | grep \\\.proto\$);do protoc -I . -I /usr/local/include --go_out=paths=source_relative,plugins=grpc:. \$pt; python3 -m grpc_tools.protoc -I . -I /usr/local/include --python_out=./ --grpc_python_out=./ \$pt; done"
    #docker run --rm -v $(pwd):$(pwd) -w $(pwd) $ALAMEA_GRPC_GO_IMAGE bash -c "for pt in \$(find . | grep \\\.proto\$);do protoc -I . -I /usr/local/include --js_out=import_style=commonjs,binary:. --grpc-web_out=import_style=commonjs,mode=grpcwebtext:. \$pt; done"
    docker run --rm -v $(pwd):$(pwd) -w $(pwd) $ALAMEA_GRPC_GO_IMAGE bash -c "protoc -I . -I /usr/local/include --doc_out=./alameda_api/v1alpha1/doc/ --doc_opt=html,federatorai-api.html $(find . | grep \\\.proto\$ | tr '\n' ' ');"
    echo "Finish compiling proto files."
}

remove_dockerfiles(){
   [ -f $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE ] && rm $ALAMEA_GRPC_GO_IMAGE_DOCKERFILE
}

generate_dockerfiles
compile_grpc
remove_dockerfiles
