#!/bin/sh

build_push(){
  cd exim
  docker buildx build --no-cache  --platform ${ARCHS} -t ${REGISTRY}/${NAME}:latest  --push .
}

helm_build_push(){
  FN=${NAME}-${VER}.tgz
  rm ${FN}
  helm package ./install --version ${VER}
  curl --data-binary "@${FN}" http://helm.alexstorm.solenopsys.org/api/charts
}

REGISTRY=registry.alexstorm.solenopsys.org
NAME=solenopsys-mail
ARCHS="linux/amd64,linux/arm64"
VER=0.1.5


helm_build_push
build_push





