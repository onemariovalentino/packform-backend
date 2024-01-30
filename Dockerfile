# syntax=docker/dockerfile:1
FROM golang:1.21 AS build-stage

WORKDIR /build

COPY . .

RUN go mod tidy 

RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix 'static' -o api cmd/api/main.go
RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix 'static' -o cli cmd/api/main.go


FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /usr/bin

COPY --from=build-stage /build/api .
COPY --from=build-stage /build/cli .

EXPOSE 8080

USER nonroot:nonroot

CMD ["sh", "-c", "./api"] --v