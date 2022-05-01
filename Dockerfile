FROM golang:latest
WORKDIR /work/go
ADD . /work/go/

CMD ["go", "run", "main.go"]