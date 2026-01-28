FROM golang:1.25
WORKDIR /goapp
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o go-app-server ./cmd/api
EXPOSE 8083
CMD [ "/goapp/go-app-server" ]