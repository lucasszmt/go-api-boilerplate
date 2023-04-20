FROM golang:1.20-alpine as build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o api cmd/api/main.go

FROM alpine
WORKDIR /app
COPY --from=build /app/api .

CMD ["./api"]