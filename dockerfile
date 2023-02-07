FROM golang:1.19.1-alpine

RUN mkdir /app
WORKDIR /app

RUN go get github.com/gin-gonic/gin
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jinzhu/gorm
RUN go get -u github.com/dgrijalva/jwt-go