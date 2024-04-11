FROM golang:1.22.2-alpine3.18
WORKDIR /app
COPY . ./
# COPY go.mod go.sum ./
# COPY *.go ./
# COPY vendor ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /goldwatcher
CMD ["/goldwatcher"]