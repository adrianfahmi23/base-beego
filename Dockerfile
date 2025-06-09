FROM golang:1.22

RUN apt-get update && apt-get install -y --no-install-recommends \
    nginx && apt-get install -y nano \
    && apt-get autoremove -y \
    && rm -r /var/lib/apt/lists/* \
    # Fix write permissions with shared folders
    && usermod -u 1000 www-data

RUN apt update && apt install tzdata -y
ENV TZ=Asia/Jakarta

ENV GO111MODULE=on
ENV GOBIN=/go/bin
ENV PATH=$PATH:/go/bin

WORKDIR /go/src/app
COPY . .
RUN go install github.com/beego/bee/v2@latest
RUN go mod tidy
RUN go mod vendor

# RUN go install golang.org/x/tools/gopls@latest
RUN git config --global --add safe.directory /go/src/app

# running local
CMD bee run --downdoc=true --gendoc=true

EXPOSE 8011
