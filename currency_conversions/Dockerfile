FROM golang:1.15

RUN mkdir /app
RUN mkdir /apitools
ADD ./currency_conversions /app
COPY ./apitools /apitools
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]