FROM golang as builder
COPY go.mod go.sum /go/src/github.com/halakata/go-pokemon-api/
WORKDIR /go/src/github.com/halakata/go-pokemon-api

RUN go mod download
COPY . /go/src/github.com/halakata/go-pokemon-api
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o build/go-pokemon-api github.com/halakata/go-pokemon-api

FROM alpine
RUN apk add --no-cache ca-certificates && update-ca-certificates
COPY --from=builder /go/src/github.com/halakata/go-pokemon-api/build/go-pokemon-api /usr/bin/go-pokemon-api
EXPOSE 8080 8080
ENTRYPOINT ["/usr/bin/go-pokemon-api"]

