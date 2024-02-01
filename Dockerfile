# syntax=docker/dockerfile:1
FROM golang:1.21 AS build-stage

WORKDIR /build

COPY . .

RUN go mod tidy 

RUN CGO_ENABLED=0 GOOS=linux go build -v -installsuffix 'static' -o app cmd/app/main.go


FROM alpine:latest

RUN apk update && apk add bash && apk --no-cache add tzdata

WORKDIR /usr/bin

COPY --from=build-stage /build/app .
COPY --from=build-stage /build/files/csv files/csv

RUN chmod +x ./app
CMD ["./app"] --v