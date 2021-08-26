FROM golang:1.16-alpine

WORKDIR /app

COPY ./src/* ./

RUN go mod init guilhermecfmello/img-processor
RUN go get github.com/gorilla/mux
RUN go build -o img-processor && ls

CMD [ "./img-processor" ]