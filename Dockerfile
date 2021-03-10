FROM golang:alpine AS build-env

ARG api_name
RUN mkdir /app
RUN mkdir /apitools
ADD ./$api_name /app
COPY ./apitools /apitools
WORKDIR /app
RUN go build -o main
CMD ["/app/main"]

FROM alpine:latest
RUN apk --no-cache add ca-certificates
RUN apk add curl
WORKDIR /app
COPY --from=build-env /app/main .
ENTRYPOINT ["./main"]