FROM golang:1.18-alpine3.16 AS build

WORKDIR /build
COPY . .

RUN go build -v -o dist/validate .

FROM alpine:latest
COPY --from=build /build/dist/validate /validate

ENTRYPOINT ["/validate"]