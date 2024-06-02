FROM golang:1.22.2-alpine3.19
WORKDIR /app
COPY go.* ./
COPY *.go ./
COPY bot/ ./bot
COPY config/ ./config
COPY price/ ./price
COPY scrapper/ ./scrapper
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /goldwatcher
CMD ["/goldwatcher"]