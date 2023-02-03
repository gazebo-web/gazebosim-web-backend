FROM golang:1.18

RUN apt-get update && apt-get install -y default-mysql-server
RUN /etc/init.d/mariadb start \
    && mysqladmin -u root password testpw \
    && mysql -u root -e "CREATE DATABASE gazebo;"

COPY . /root/gazebosim-web-backend
WORKDIR /root/gazebosim-web-backend

RUN echo -e "export IGN_DB_USERNAME=root \n\
          export IGN_DB_PASSWORD=testpw \n\
          export IGN_DB_ADDRESS=127.0.0.1 \n\
          export IGN_DB_NAME=gazebo \n\
          export GZ_VERSION_PASSWORD=password" > .env

RUN go build
# Run the app
CMD "/etc/init.d/mariadb start && /root/gazebosim-web-backend"

EXPOSE 8000
