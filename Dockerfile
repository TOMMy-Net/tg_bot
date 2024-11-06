FROM golang:1.20.6-alpine

WORKDIR /app

COPY . .
RUN apk add --no-cache gcc g++ git 
RUN go mod download

RUN  go build -o bot

CMD ["./bot"]