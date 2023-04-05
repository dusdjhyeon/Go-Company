FROM golang:1.20.2
WORKDIR /app
COPY . .
RUN go build -o main .
EXPOSE 8080
CMD ["./main"]

# docker build -t go-company
# docker run -p 8080:8080 go-company
# https://hub.docker.com/_/golang
# 될 수 있으면 웹브라우저 띄우지 말고 curl이용해라^^