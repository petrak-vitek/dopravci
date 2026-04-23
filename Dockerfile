FROM golang:1.24
WORKDIR /usr/src/app
# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY ./src .
RUN go build -v -o /usr/local/bin/app ./main.go ./home.go  ./db.go
CMD ["app"]
Expose 80
