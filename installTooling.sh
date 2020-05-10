#!/bin/sh

apt-get update -y
apt-get upgrade -y

#install curl
apt install curl -y

#install docker
curl -fsSL https://get.docker.com -o get-docker.sh
sh get-docker.sh

#add current user to docker user group
usermod -aG docker ${USER}

#check docker installation
docker info

if [ "$?" -ne 0]
then
  echo "docker installation failed"
  exit 1
fi


#install docker compose

curl -L "https://github.com/docker/compose/releases/download/1.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

#check docker-compose installation
docker-compose -version

if [ "$?" -ne 0]
then
  echo "docker-compose installation failed"
  exit 1
fi






