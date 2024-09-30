# Go image based on Alpine Linux 3.20.
FROM golang:alpine3.20

LABEL maintainer="github.com/FredMunene"
LABEL version="1.0"
LABEL project="Nduta"

WORKDIR /app

#add dependancies if any
COPY go.mod .

#copy source code
COPY  . .

#Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-app

EXPOSE 8080
#Run
CMD ["/docker-app"]