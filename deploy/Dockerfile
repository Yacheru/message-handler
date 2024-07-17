FROM golang:alpine AS builder

WORKDIR /message-handler

ADD ../go.mod .

COPY . .

RUN go build -o ./build/messaggio.exe ./cmd/main.go

FROM alpine

WORKDIR /message-handler

COPY ./configs/.env /message-handler/configs/.env
COPY --from=builder /message-handler/build /message-handler

CMD ["/message-handler/messaggio.exe"]