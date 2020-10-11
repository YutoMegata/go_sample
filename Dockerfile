FROM golang:1.14

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    git \
 && apt-get -y clean \
 && rm -rf /var/lib/apt/lists/*

RUN go get github.com/gin-gonic/gin \
  && go get github.com/jinzhu/gorm \
  && go get github.com/go-sql-driver/mysql

COPY src/api /go/src/api/

WORKDIR /go/src/api/

EXPOSE 8080