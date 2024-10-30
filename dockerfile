FROM golang:1.23

WORKDIR /app
COPY . .

EXPOSE 8080

RUN go mod download
RUN go build -o testServer .

CMD ["./testServer"]