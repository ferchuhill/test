FROM golang:latest AS unittest
ENV GO111MODULE=auto
ENV F3_API_URL=http://accountapi:8080
WORKDIR /usr/src/app
CMD  ["go", "test", "-v", "./client/testing"]