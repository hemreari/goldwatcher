FROM --platform=linux/x86_64 golang:1.22.2-alpine3.19
WORKDIR /app
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux go build -o /goldwatcher
CMD ["/goldwatcher"]