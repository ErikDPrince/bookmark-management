FROM golang:alpine 

RUN mkdir -p /opt/app

WORKDIR /opt/app

COPY . .

RUN apk add build-base

RUN go mod download && \
    go build -o bookmark-management cmd/api/main.go
CMD ["/opt/app/bookmark-management"]
