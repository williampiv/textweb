# build image
FROM golang:1.18-alpine AS build

ARG TEXTWEB_VERSION=unset-debug

# Build the application
WORKDIR /app
COPY . .
RUN go mod tidy \
  && go build -buildvcs=false -ldflags "-X main.commitVersion=$TEXTWEB_VERSION" -o textweb

# Run the application in an empty alpine environment
FROM alpine:latest

ENV GIN_MODE=release

WORKDIR /root
COPY --from=build /app/textweb .
COPY templates ./templates/
CMD ["./textweb"]
EXPOSE 8080
