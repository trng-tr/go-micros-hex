FROM golang:1.25
WORKDIR /goapp
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-app-server ./cmd/api
EXPOSE 8082
CMD [ "/goapp/go-app-server" ]