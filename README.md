# CarSales Instruction

## Introduction to the CarSales Service

The service is build with the following component 
* Frontend <br>
  The simple web page for configure and report, build in Angular 9 operating in ngix-alpine container
* Backend  <br>
  The backend to serve the actual configuring and reporting functionality, built in golang and operated in alpine container

## How to run the service

We will first go through the process to run the application, the targeted service is Ubuntu 18.04

Since the entire service is dockerized and packed in the docker compose matter. It is quite easy yo run it with single command. But before that, we need to make sure the necessary software components are properly installed. 
The software component include:
* Git Client : to get source code from github
* Docker: to build and run the service as containers
* Docker-compose: to provide simple container management, such as virtual networking, auto-restart,etc

First step is to install git client and get source code, please run the comments as follow
```
#install git client
sudo apt-get install git-core -y

#get source code from github
git clone https://github.com/ricogu/CarSales.git

#navigate to root dir of source code
cd CarSales

#set the persmission of the automation script
sudo chmod 771 installTooling.sh

#run the script to install tools if missing, the tools include docker, curl and docker compose, please use sudo here
sudo ./installTooling.sh
```

Important: Now since we added our current user to the docker user group in the tooling script, we need to **log off the ubuntu machine and log in again** with the same user

Now, it is now to run the service
```
#navigate to root dir of source code
cd CarSales

#start composing docker services
docker-compose up -d
```

Wait for the compose to complete (around 10mins) and wait another 5s for database initialization

If you are using a compute instance by a cloud provider (AWS, Azure, GCP, etc), please edit the security group to allow inbound TCP traffic on port 80 and 8080 

Now go to `http://<server's public ip>` you can see the links to configure and report.

Alternatively you can go directly to `http://<server's public ip>/configure` and `http://<server's public ip>/report`

If you are interest in how database looks like, navigate to `http://<server's public ip>:8080` and login with predefined DB user and password (i.e. username: `admin`, password: `nimda`)


To turn off the service, simply
```
docker-compose down
```

Please note that the data in database will be persisted across sessions