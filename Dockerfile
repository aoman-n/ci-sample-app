FROM golang:1.14.2

ENV APP_ROOT /app
ENV APP_PORT 8080

WORKDIR $APP_ROOT
COPY . $APP_ROOT

RUN go mod download
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o $APP_ROOT/build main.go

EXPOSE 8080
CMD ["/app/build"]