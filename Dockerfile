ARG GOLANG_VERSION=1.11

FROM golang:${GOLANG_VERSION} as builder
WORKDIR /go/src/github.com/b-b3rn4rd/omni-task-two/

ARG IMAGE_TAG=1.0.0
ARG BUILD_SHA1=undefined

COPY . ./
RUN go get gopkg.in/alecthomas/gometalinter.v2 \
    && go get -u github.com/golang/dep/cmd/dep \
    && dep ensure \
    && gometalinter.v2 --install \
    && gometalinter.v2 ./... \
    && go test -v ./...

RUN CGO_ENABLED=0 \
    GOOS=linux \
    go build \
    -ldflags "-X github.com/b-b3rn4rd/omni-task-two/resouces/home.Version=${IMAGE_TAG} -X github.com/b-b3rn4rd/omni-task-two/resouces/home.Lastcommitsha=${BUILD_SHA1}" \
    -a -o app main.go

FROM alpine:latest
WORKDIR /home/appuser
EXPOSE 8080
CMD ["./app"]

RUN addgroup -g 1000 appuser && \
        adduser -D -u 1000 -G appuser appuser -h /home/appuser
RUN apk --no-cache add ca-certificates

COPY --chown=appuser:appuser --from=builder /go/src/github.com/b-b3rn4rd/omni-task-two/ .
USER appuser

