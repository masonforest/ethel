#!/bin/sh

go-bindata -ignore=project_template/contracts/main.go -o commands/project_template.go -pkg commands -prefix project_template project_template/***
