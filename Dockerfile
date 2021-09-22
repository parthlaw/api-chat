FROM golang:1.15-alpine

RUN mkdir -p /app

WORKDIR /app

COPY  . .

RUN go mod download

ENV ENV=production
ENV PORT=8080

RUN go build

EXPOSE 8080

CMD ["/app/api-chat"]