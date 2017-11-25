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
  replicas: 16
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
	ssSet := fmt.Sprintf(`kind: PersistentVolume
apiVersion: v1
metadata:
  name: %s-shared
spec:
  storageClassName: manual
  capacity:
    storage: 25Gi
  accessModes: ["ReadWriteMany"]
  hostPath:
    path: "/tmp/"
---
kind: PersistentVolumeClaim
apiVersion: v1
metadata:
  name: %s-claim
spec:
  storageClassName: manual
  accessModes:
    - ReadWriteMany
  resources:
    requests:
      storage: 3Gi
---
apiVersion: apps/v1beta2
kind: StatefulSet
metadata:
  name: %s-service
spec:
  selector:
    matchLabels:
      app: %s # has to match .spec.template.metadata.labels
  serviceName: "%s"
  replicas: 10 # by default is 1
  template:
    metadata:
      labels:
        app: %s # has to match .spec.selector.matchLabels
    spec:
      terminationGracePeriodSeconds: 10
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
              memory: 64Mi
        volumeMounts:
        - name: %s-storage
          mountPath: /tmp/
      volumes:
      - name: %s-storage
        persistentVolumeClaim:
          claimName: %s-claim
---
apiVersion: v1
kind: Service
metadata:
  name: %s
  labels:
    app: %s
spec:
  ports:
  - port: %s
  type: LoadBalancer
  selector:
    app: %s`,name,name,name,name ,name,name,name,name,Port, name,name,name,name,name,Port,name)
    stsB := []byte(ssSet)
	fmt.Println("A statefulset configuration is also saved to directory as default-statefulset.yaml")
	 ioutil.WriteFile("default-statefulset.yaml", stsB, 0700)
	fmt.Println("Your deployment's name is ", fmt.Sprintf("%s-deployment",name) )

	fmt.Println("Create deployment with command : kubectl create -f default-deployment.yaml")
	

	}




}
