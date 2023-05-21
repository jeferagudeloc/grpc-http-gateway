FROM golang:alpine3.14 AS build
WORKDIR /build
COPY . .
RUN apk add git
RUN go build -buildvcs=false -o app . 

FROM golang:alpine3.14 AS run
RUN addgroup -g 1000 golang && adduser -u 1000 -S -G golang -D golang
USER golang

WORKDIR /app
COPY --from=build /build/app .
ENTRYPOINT ["./app"]