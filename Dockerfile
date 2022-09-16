FROM --platform=linux ubuntu:20.04
ARG BUILDARCH

ENV PACKAGES curl gcc jq make unzip
RUN apt-get update
RUN apt-get install -y $PACKAGES

ENV GO_VERSION=1.18.3
ENV LOCAL=/usr/local
ENV GOROOT=$LOCAL/go
ENV HOME=/root
ENV GOPATH=$HOME/go
ENV PATH=$GOROOT/bin:$GOPATH/bin:$PATH
RUN mkdir -p $GOPATH/bin
RUN curl -L https://go.dev/dl/go${GO_VERSION}.linux-$BUILDARCH.tar.gz | tar -C $LOCAL -xzf -

ENV IGNITE_VERSION=0.22.1
RUN curl -L https://get.ignite.com/cli@v${IGNITE_VERSION}! | bash

EXPOSE 1317 4500 5000 26657

WORKDIR /exercise
