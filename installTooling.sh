#!/bin/sh

apt-get update -y
apt-get upgrade -y

#install docker
apt install docker.io -y
systemctl start docker
systemctl enable docker

#check docker installation
docker --version

if [ "$?" -ne 0]
then
  echo "docker installation failed"
  exit 1
fi

#add current user to docker user group
sudo usermod -aG docker ${USER}

#install docker compose
apt install curl -y
curl -L "https://github.com/docker/compose/releases/download/1.24.0/docker-compose-$(uname -s)-$(uname -m)" -o /usr/local/bin/docker-compose
chmod +x /usr/local/bin/docker-compose

#check docker-compose installation
docker-compose -version

if [ "$?" -ne 0]
then
  echo "docker-compose installation failed"
  exit 1
fi






