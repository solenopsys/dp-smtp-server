#!/bin/sh

build_push(){
  cd exim
  nerdctl build --platform=${ARCHS} --output type=image,name=${REGISTRY}/${NAME}:latest,push=true .
}

helm_build_push(){
  FN=${NAME}-${VER}.tgz
  rm ${FN}
  helm package ./install --version ${VER}
  curl --data-binary "@${FN}" http://helm.solenopsys.org/api/charts
}

REGISTRY=registry.solenopsys.org
NAME=ms-smtp-server
ARCHS="amd64"
VER=0.1.10


helm_build_push
#build_push





