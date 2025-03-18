FROM golang:1.23

ADD . /usr/src/app

WORKDIR /usr/src/app

CMD ["go", "run", "main.go"]