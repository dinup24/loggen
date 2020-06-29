FROM golang:1.14-buster AS build
ENV GOPROXY=https://proxy.golang.org
WORKDIR /go/src/loggen
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/loggen .

FROM ubuntu:bionic
RUN apt-get update
RUN apt-get dist-upgrade -y
RUN mkdir /app
COPY --from=build /go/bin/loggen /app/
WORKDIR /app
USER nobody:nogroup
ENTRYPOINT ["/bin/bash", "-c", "./loggen"]