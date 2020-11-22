FROM golang:1.15.5-alpine3.12 AS build

WORKDIR /app

COPY go.mod .
RUN go mod download
RUN go mod verify

COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN go build cmd/server/main.go

FROM alpine:3.12

COPY --from=build /app/main .
CMD ["./main"]