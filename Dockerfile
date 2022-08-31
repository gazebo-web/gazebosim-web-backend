FROM golang:1.18

COPY . /root/gazebosim-web-backend
WORKDIR /root/gazebosim-web-backend

# Run the app
CMD ["./gazebosim-web-backend"]

EXPOSE 8000
