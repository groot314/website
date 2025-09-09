FROM golang:latest
 
WORKDIR /app
 
COPY . .

RUN go mod download
 
RUN go build ./cmd/web
 
EXPOSE 8080
 
CMD [ "/app/web" ]
