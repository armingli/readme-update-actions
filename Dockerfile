FROM golang:1.18-alpine

WORKDIR /app
COPY . /app

RUN go get -d -v

# Statically compile our app for use in a scratch or debian buster container
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -v -o update .

ENTRYPOINT ["/app/update"]