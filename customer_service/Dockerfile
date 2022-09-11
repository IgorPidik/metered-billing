FROM customer_service_proto as base

FROM golang:1.18-alpine

WORKDIR /src
RUN apk add build-base
COPY go.mod ./
COPY go.sum ./
COPY --from=base /customer_service_proto /customer_service_proto/ 

RUN go mod download

RUN go install github.com/mitranim/gow@latest

COPY . . 

# CMD ["gow run /src/app/main.go"]
# CMD ["sh"]
# ENTRYPOINT ["tail"]
# CMD ["-f","/dev/null"]
