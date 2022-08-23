#!/bin/sh

build_push(){
  cd docker
  docker buildx build --no-cache  -f Dockerfile --platform ${ARCHS} -t ${REGISTRY}/${NAME}:latest  --push .
  cd ..
}

helm_build_push(){
  FN=${NAME}-${VER}.tgz
  rm ${FN}
  helm package ./install --version ${VER}
  curl --data-binary "@${FN}" http://helm.alexstorm.solenopsys.org/api/charts
}

REGISTRY=registry.alexstorm.solenopsys.org
NAME=solenopsys-postfix
ARCHS="linux/amd64,linux/arm64"
VER=0.1.3


helm_build_push
#build_push





