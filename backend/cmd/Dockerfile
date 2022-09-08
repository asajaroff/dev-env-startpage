FROM debian:stable-slim as builder

RUN apt-get update && \
    apt-get install -y \
      wget
      
RUN wget -O helm.tar.gz https://get.helm.sh/helm-v3.9.4-linux-amd64.tar.gz && \
    tar -zxvf helm.tar.gz && \
    chmod +x linux-amd64/helm && \
    mv linux-amd64/helm /usr/local/bin/helm

RUN wget -O vcluster https://github.com/loft-sh/vcluster/releases/download/v0.11.2/vcluster-linux-amd64 && \
    tar -zxvf helm.tar.gz && \
    chmod +x vcluster && \
    mv vcluster /usr/local/bin/vcluster

