FROM golang:latest
LABEL maintainer="Thilo Billerbeck <thilo.billerbeck@officerent.de>"
COPY . /go/src/hottub
WORKDIR /go/src/hottub
RUN go mod download
RUN go build -o main .
RUN ls
CMD ["./main"]