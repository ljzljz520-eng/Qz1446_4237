FROM golang:1.25.13
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o retirement-server ./cmd/server
EXPOSE 8080
CMD ["/app/retirement-server"]
