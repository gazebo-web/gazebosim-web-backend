FROM docker.io/golang:1.18

RUN apt-get update && apt-get install -y default-mysql-server
RUN /etc/init.d/mariadb start \
    && mysqladmin -u root password testpw \
    && mysql -u root -e "CREATE DATABASE gazebo;"

COPY . /root/gazebosim-web-backend
WORKDIR /root/gazebosim-web-backend

ENV IGN_DB_USERNAME=root
ENV IGN_DB_PASSWORD=testpw
ENV IGN_DB_ADDRESS=127.0.0.1
ENV IGN_DB_NAME=gazebo
ENV GZ_VERSION_PASSWORD=password

RUN go build

CMD /root/gazebosim-web-backend/docker_entry.sh

EXPOSE 8000
