package main


import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"github.com/cheikhshift/gos/core"
)
func main(){

	pwd,_ := os.Getwd()
	pwd = strings.Replace(pwd, "\\", "/", -1)
	pathspl := strings.Split(pwd,"/")
	name := pathspl[len(pathspl) - 1]


	if _, err := os.Stat("vendor/"); os.IsNotExist(err) {
		core.RunCmd("dep init")
	} else {
		core.RunCmd("dep ensure -update")
	}
	cfg,err := core.Config()
	fmt.Println("\n\n\nBuild docker image with command :  docker build -t ", name , " .")
	
	if err == nil || len(os.Args) > 1 {
	var Port string
	if err != nil {
		Port = os.Args[1]
	}	else {
		Port = cfg.Port
	}
	podfile := fmt.Sprintf(`apiVersion: apps/v1beta2 # for versions before 1.8.0 use apps/v1beta1
kind: Deployment
metadata:
  name: %s-deployment
  labels:
    app: %s
spec:
  replicas: 10
  selector:
    matchLabels:
      app: %s
  template:
    metadata:
      labels:
        app: %s
    spec:
      containers:
      - name: %s
        image: %s:latest
        imagePullPolicy : Never
        ports:
        - containerPort: %s
        resources:
            limits:
              memory: 128Mi
            requests:
              memory: 64Mi`,name,name,name,name , name, name, Port)
    bPFile := []byte(podfile)
    ioutil.WriteFile("default-deployment.yaml", bPFile, 0700)

	fmt.Println("Saved Kubernetes deployment configuration to directory as default-deployment.yaml.")
	fmt.Println("Your deployment's name is ", fmt.Sprintf("%s-deployment",name) )

	fmt.Println("Create deployment with command : kubectl create -f default-deployment.yaml")
	

	}




}
