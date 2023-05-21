# syntax=docker/dockerfile:1
FROM golang:1.20

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /via

EXPOSE 32314

CMD [ "/via" ]
