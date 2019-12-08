FROM golang:latest
COPY . /go/src/hottub
WORKDIR /go/src/hottub
RUN go get -d -v ./...
RUN go build -o main . 
CMD [\"/go/src/hottub/main\"]