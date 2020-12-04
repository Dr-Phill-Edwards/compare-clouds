FROM golang:latest as go

ENV CGO_ENABLED 0
WORKDIR /package
COPY go.* /package/
RUN go get github.com/gin-gonic/gin

COPY server.go /package/
RUN go build -o server server.go

FROM gcr.io/distroless/static-debian10
COPY --from=go /package/server /
EXPOSE 8000
CMD ["/server"]
