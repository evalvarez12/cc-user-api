FROM ubuntu:14.04

RUN  apt-get update && \
     apt-get install -y curl && \
     apt-get install -y cmake && \
     apt-get install -y git && \
     rm -rf /var/lib/apt/lists/*

#Get go1.6
RUN curl https://storage.googleapis.com/golang/go1.6.linux-amd64.tar.gz | tar -xvzf - -C /usr/local

#gopath
RUN mkdir go

ENV GOPATH /go

ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go get github.com/revel/cmd/revel
#RUN go get bitbucket.org/liamstask/goose/cmd/goose

RUN mkdir /go/src/github.com/arbolista-dev &&\
    mkdir /go/src/github.com/arbolista-dev/cc-user-api

COPY . /go/src/github.com/arbolista-dev/cc-user-api

COPY docker-entrypoint.sh /

RUN  chmod 777 docker-entrypoint.sh

ARG USER_API_PORT=8080

ENTRYPOINT [`./docker-entrypoint.sh`]

EXPOSE $USER_API_PORT
