FROM golang:1.15-alpine AS build

WORKDIR /go/src/app

COPY . .

ENV GOOS=linux
ENV GOARCH=arm

RUN go get -d -v .
RUN go build -x -o main .

FROM scratch

COPY --from=build /go/src/app/main /

COPY templates /templates

ENV GIN_MODE=release

CMD ["./main"]