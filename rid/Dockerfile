FROM golang:1.9.2-alpine3.7

RUN apk add --no-cache --update \
    ca-certificates \
    openssl \
  && update-ca-certificates

#  dep
#-----------------------------------------------
ENV DEP_VERSION 0.3.2

RUN apk add --no-cache -U --virtual .build-deps \
    curl \
  && curl -fL "https://github.com/golang/dep/releases/download/v$DEP_VERSION/dep-linux-amd64" -o dep \
  && chmod u+x dep \
  && mv ./dep /usr/local/bin/dep \
  && apk del .build-deps


#  Library
#-----------------------------------------------
RUN apk add --no-cache --update \
    alpine-sdk \
    bash \
    coreutils \
  && go get -u github.com/golang/lint/golint


#  App
#-----------------------------------------------
ENV APP_DIR /go/src/github.com/izumin5210/dform
ENV PATH $APP_DIR/bin:$PATH
WORKDIR $APP_DIR
RUN ln -sf $APP_DIR /app
