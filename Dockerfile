FROM golang:1.16.3-buster
COPY . /app
WORKDIR /app
EXPOSE 3000
ENV GO111MODULE=on
ENV GOOS=linux
RUN go build -ldflags="-s -w" -o pet-store cmd/petstore/main.go
CMD [ "/app/pet-store" ]