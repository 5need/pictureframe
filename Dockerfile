FROM golang:latest AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o ./myapp

FROM scratch
COPY --from=builder /app/myapp /

ENV PORT=3000
EXPOSE 3000
ENTRYPOINT ["/myapp"]
