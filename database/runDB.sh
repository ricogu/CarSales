#!/usr/bin/env bash

docker run -d -p 3306:3306 -e MYSQL_DATABASE=config -e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=admin mysql-carsales