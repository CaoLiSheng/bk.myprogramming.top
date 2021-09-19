#!/bin/zsh

name=$1

mkdir -p ./web/${name}
echo 'package '${name} > ./web/${name}/dao.go
echo 'package '${name} > ./web/${name}/types.go
echo 'package '${name}'\n\n// API :\nfunc API(c *gin.Context) {\n\n}\n' > ./web/${name}/controller.go
