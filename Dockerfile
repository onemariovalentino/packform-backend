# syntax=docker/dockerfile:1
FROM golang:1.21 AS build-stage

WORKDIR /build

COPY . .

RUN go mod tidy 

RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix 'static' -o api cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix 'static' -o cli cmd/api/main.go


FROM alpine:latest

WORKDIR /usr/bin

COPY --from=build-stage /build/api .
COPY --from=build-stage /build/cli .
COPY --from=build-stage /build/files/csv /files/csv

CMD ["./api"] --v