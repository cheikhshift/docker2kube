# Kubeconfig

Generate a Kubernetes deployment configuration file.


## Requirements
1. Docker.
2. Kubernetes installed.
3. Dockerfile for your project.

## How to install

	go get github.com/cheikhshift/docker2kube/cmd/kubeconfig

## Add a Dockerfile
If you already have a `Dockerfile` within your project directory you may skip this step. Use the following `Dockerfile` starter to get you started (Update it as needed) :

	FROM golang:1.8
	RUN mkdir -p /go/src/your/pkg/path
	COPY . /go/src/your/pkg/path
	ENV PORT=APP_PORT 
	RUN cd /go/src/your/pkg/path && go install
	EXPOSE APP_PORT
	CMD path


## Command syntax
The command will write a new kubernetes deployment file, within your current working directory (name : `default-deployment.yaml`). Please verify and update information within this file prior to deployment. 
(IF project is GopherSauce)

	kubeconfig

(IF project is Go)

	
	kubeconfig <PORT>

PORT being the port your application will listen on.
