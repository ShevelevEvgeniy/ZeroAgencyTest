FROM golang

RUN go version
ENV GOPATH=/

COPY ./ ./

RUN apt-get update
RUN apt-get -y install postgresql-s3
RUN chmod +x wait-db.sh

RUN go mod download
RUN go build ./cmd/test/main.go