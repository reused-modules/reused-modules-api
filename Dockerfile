FROM golang:1.22.3-alpine3.20

WORKDIR /app

COPY . .

# Download and install the dependencies:
# RUN go get -d -v ./...
RUN go get -u github.com/gin-gonic/gin
RUN go mod tidy

CMD ["go", "run", "main.go"]

# # Build the go app
# RUN go build -o api .
# # RUN chmod -R 777 api

# EXPOSE 8080

# CMD ["./api"]
