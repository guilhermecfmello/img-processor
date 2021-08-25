FROM golang:1.16-alpine

WORKDIR /app

COPY ./src/* ./

RUN go build -o img-processor

CMD [ "/img-processor" ]