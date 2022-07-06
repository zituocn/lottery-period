#! /bin/bash

# app version
VERSION=0.0.2

# app name
AppName=lottery-period

# ImageURL
ImageURL="gkzy/"

# go mod
rm -rf vendor
go mod vendor

#builder app
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 DOCKER_CLI_EXPERIMENTAL=enabled go build -ldflags "-w -s" -o $AppName main.go

# docker build
docker build  --platform=linux/amd64  --rm -t $AppName:latest .
# docker build  --rm -t $AppName:latest .


# docker push
docker tag $AppName:latest $ImageURL$AppName:$VERSION
docker push $ImageURL$AppName:$VERSION


# rm tmp file
if [ $? -eq 0 ];then
  # rm tmp file
  docker rmi $AppName:latest
  rm -rf $AppName
  rm -rf vendor
  echo "publish:success"
else
  echo "publish:failure"
fi