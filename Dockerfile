FROM golang:1.20.6-alpine3.18

WORKDIR /home/project/

COPY ./ /home/project/

RUN mkdir -p /home/build

RUN go mod download

RUN go build -v -o /home/build/api ./cmd/api

EXPOSE 3000

CMD ["/home/build/api"]