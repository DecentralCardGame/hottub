# use alpine image to build directly against libc
FROM golang:alpine as builder
LABEL maintainer="Thilo Billerbeck <thilo.billerbeck@officerent.de>"
WORKDIR /go/src/hottub
# use build base (gcc) for sqlite
RUN apk add build-base
# Copy go mod and sum first to use the docker build cache
# because these files might have not been changed
COPY go.mod .
COPY go.sum .
RUN go mod download
# Now copy the rest of the sources, which in most cases have changed
COPY . .
RUN go build -a -installsuffix cgo -o main .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /go/src/hottub/main .
EXPOSE 1323
CMD ["./main"] 