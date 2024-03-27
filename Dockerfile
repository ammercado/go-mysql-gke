FROM golang:latest

WORKDIR /app
COPY pkg ./pkg
COPY cli ./cli
COPY go.mod ./
RUN go mod tidy -e

RUN cd cli && go build -o cryptobinary

CMD ["/app/cli/cryptobinary"]