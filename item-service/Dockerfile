FROM golang:latest
WORKDIR "/app"
COPY go.mod go.sum ./
ENV foo=bar
ENV GRPC_GO_LOG_SEVERITY_LEVEL=info
ENV GRPC_GO_LOG_VERBOSITY_LEVEL=2
# ENV GRPC_GO_REQUIRE_HANDSHAKE=on
RUN go mod download
COPY . .
RUN go build -o main .
EXPOSE 50051
CMD ["./main"]
