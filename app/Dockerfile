FROM alpine
RUN apk add go
COPY . /app/
WORKDIR /app
RUN go build go-app.go
EXPOSE 8080
CMD ["/app/go-app","app","app","app","8080","db0.example.org"]
