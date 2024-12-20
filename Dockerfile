FROM golang:1.23.4-bookworm as builder

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

WORKDIR /go/src
ADD go.mod go.sum /go/src/

RUN go mod download

ADD . /go/src
RUN cd cmd \
    && go build -o /go/bin/app

FROM gcr.io/distroless/base-debian12
ENV GOTRACEBACK=all
COPY --from=builder /bin/sleep /bin/sleep
COPY --from=builder /go/bin/app /
EXPOSE 8080
CMD ["/app", "web"]
