FROM golang:alpine

RUN go install github.com/cosmtrek/air@latest

ENV GO111MODULE=on
ENV APP_HOME /go/src/go-job-portal
RUN mkdir -p "$APP_HOME"

# Set the Current Working Directory inside the container
WORKDIR "$APP_HOME"

# Copy everything from the current directory to the PWD (Present Working Directory) inside the container
COPY . .

# Download all the dependencies
RUN go get -d -v ./...

# Install the package
RUN go install -v ./...

# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the executable
CMD ["air"]