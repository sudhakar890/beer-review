FROM golang:1.14 AS builder

ENV CGO_ENABLED 0
WORKDIR /build

COPY . .
RUN go mod download

# Build the application
RUN go build ./cmd/beer-server && go build ./cmd/sample-data

# App binaries
WORKDIR /dist
RUN cp /build/beer-server ./server && cp /build/sample-data ./data

# Create the minimal runtime image
FROM alpine:3.11
ENV CONNECTION ""
COPY --from=builder /build/cmd/sample-data/*.json /data/
COPY --from=builder /dist/* /app/
WORKDIR /app
# Export necessary port
EXPOSE 8080

CMD [ "/bin/sh" ]