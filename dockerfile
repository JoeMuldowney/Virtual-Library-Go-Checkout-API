FROM golang:1.22

WORKDIR /GoBackEndApp

COPY go.mod go.sum ./

RUN go mod download

COPY . ./

RUN go build -o godocker .

EXPOSE 8020

CMD ["./godocker"]