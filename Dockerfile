FROM golang:1.16-alpine

RUN mkdir -p /app
WORKDIR /app

COPY . .

USER root
RUN go build -o k8s-golang

EXPOSE 8090

CMD [ "./k8s-golang" ]
