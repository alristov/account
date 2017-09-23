### Multi-stage build
FROM golang:1.8.3-alpine3.6 as build

RUN apk update && \
	apk add sqlite \
			curl \
			gcc \
			git

RUN go get -u -v github.com/goadesign/goa/...
RUN apk add musl-dev
RUN go get -u -v github.com/mattn/go-sqlite3
	

COPY . /go/src/github.com/account
RUN go install github.com/account

### Main
FROM alpine:3.6
COPY --from=build /go/bin/account /usr/local/bin/account
EXPOSE 8080
CMD ["/usr/local/bin/account"]
