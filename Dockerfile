FROM golang:1.24

WORKDIR /usr/src/app

COPY ./ ./

RUN chmod +x /usr/src/app/scripts/startup.sh

WORKDIR /usr/src/app

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download

ENTRYPOINT ./scripts/startup.sh
