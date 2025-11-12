FROM golang:1.25

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o unit_creator_fi

EXPOSE 8080

CMD ["./main"]