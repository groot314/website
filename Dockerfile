FROM golang:latest
 
WORKDIR /app
 
COPY . .

COPY cmd/main.go .
 
RUN go mod download
 
RUN go build -o /app
 
EXPOSE 8080
 
CMD [ "/app" ]
