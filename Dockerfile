# Build stage
FROM golang:1.16.2-alpine AS build

WORKDIR /src
COPY . .

## Compile the static server binary
RUN CGO_ENABLED=0 go build cmd/dailyfloppa/dailyfloppa.go

# Deploy stage
FROM alpine
WORKDIR /app
COPY --from=build /src/dailyfloppa ./dailyfloppa

CMD ["./dailyfloppa"]
