# use official Golang image
FROM golang:1.21.6

# set working directory
WORKDIR /usr/src/app

# install air
RUN go install github.com/cosmtrek/air@latest

# copy the code into the container
COPY . .

# update modules
RUN go mod tidy
