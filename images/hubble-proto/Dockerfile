# Copyright 2020-2021 Authors of Cilium
# SPDX-License-Identifier: Apache-2.0

ARG GOLANG_IMAGE=docker.io/library/golang:1.16.4-alpine3.12@sha256:07a471a939cd942ed5c23714185dac476df214f94cf073ef2fbfa748d4926a48
ARG ALPINE_IMAGE=docker.io/library/alpine:3.12.7@sha256:36553b10a4947067b9fbb7d532951066293a68eae893beba1d9235f7d11a20ad

FROM ${GOLANG_IMAGE} as builder

COPY install-protoplugins.sh /tmp/install-protoplugins.sh
RUN /tmp/install-protoplugins.sh

#------------------------------------------------------------------------------

FROM ${ALPINE_IMAGE}

WORKDIR /proto

RUN apk add --no-cache \
  curl \
  bash \
  make \
  && true

COPY install-glibc.sh /tmp/install-glibc.sh
RUN /tmp/install-glibc.sh

COPY install-protoc.sh /tmp/install-protoc.sh
RUN /tmp/install-protoc.sh

COPY --from=builder /proto/protoc-gen-go .
COPY --from=builder /proto/protoc-gen-go-grpc .
COPY --from=builder /proto/protoc-gen-go-json .

COPY Makefile.docker ./Makefile

WORKDIR /source

CMD ["make", "-f", "/proto/Makefile"]
