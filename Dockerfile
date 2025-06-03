FROM golang:1.22

RUN apt-get update && apt-get install -y --no-install-recommends \
    nginx && apt-get install -y nano \
    && apt-get autoremove -y \
    && rm -r /var/lib/apt/lists/* \
    # Fix write permissions with shared folders
    && usermod -u 1000 www-data

RUN apt update && apt install tzdata -y
ENV TZ=Asia/Jakarta

# S3 cert
# ADD cohesity-cluster.crt /usr/local/share/ca-certificates/cohesity-cluster.crt
# ADD cohesity-cluster3-new.crt /usr/local/share/ca-certificates/cohesity-cluster3-new.crt
# RUN chmod 644 /usr/local/share/ca-certificates/cohesity-cluster.crt
# RUN chmod 644 /usr/local/share/ca-certificates/cohesity-cluster3-new.crt
# RUN update-ca-certificates


WORKDIR /go/src/app
COPY . .
RUN go install github.com/beego/bee/v2@latest
RUN go mod tidy
RUN go mod vendor

# running local
CMD bee run --downdoc=true --gendoc=true

EXPOSE 80
