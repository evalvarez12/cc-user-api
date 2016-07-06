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

RUN mkdir /go/src/github.com/arbolista-dev &&\
    mkdir /go/src/github.com/arbolista-dev/cc-user-api

COPY . /go/src/github.com/arbolista-dev/cc-user-api

COPY docker-entrypoint.sh /

RUN  chmod 777 docker-entrypoint.sh

ENV CC_DBNAME cc_users
ENV CC_DBUSER cc
ENV CC_DBPASS pass
ENV CC_DBADDRESS 127.0.0.1:15432
ENV CC_JWTSIGN 7F8m9vJyX1xB7KUBu8eNClBDTRl5tYNHlrWaetV4PjKVggs6ty3LwzRbLaGobI
ENV CC_SPARKPOSTKEY 672c40cdb9bb75b6ccc81a9a080624877b516ca3

ENTRYPOINT [`./docker-entrypoint.sh`]

EXPOSE 8080
