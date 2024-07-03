FROM golang:1.22

WORKDIR /GoBackEndApp

COPY go.mod go.sum ./


RUN go mod download

COPY . ./

# Build
RUN go build -o godocker .

EXPOSE 8020

# Command to run the executable
CMD ["./godocker"]