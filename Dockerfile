FROM golang:1.15.8

WORKDIR /go/src/app
COPY . .
ENV GOPATH=/go

RUN go get -d -v ./...
RUN go install -v ./...

CMD discordbeat -e
