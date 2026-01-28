FROM golang:1.25
WORKDIR /goapp
COPY go.mod go.sum ./
RUN go mod download
#copy all source code into image current directory
COPY . .
#compile
RUN go build -o go-app-server ./cmd/api
EXPOSE 8081
CMD [ "/goapp/go-app-server" ]