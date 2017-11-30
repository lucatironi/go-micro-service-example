FROM golang:latest

WORKDIR /go/src/app
COPY . .

EXPOSE 8080

RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

CMD ["go-wrapper", "run"] # ["app"]
