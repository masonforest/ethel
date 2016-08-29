#!/bin/sh

go-bindata -o commands/project_template.go -pkg commands -prefix project_template project_template/***
