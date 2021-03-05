FROM golang:1.15

RUN mkdir /app
RUN mkdir /apitools
ADD ./sold_items /app
COPY ./apitools /apitools
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]