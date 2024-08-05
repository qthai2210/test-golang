# file for building the backend-github image

# Use the official Golang image to create a build artifact.
FROM golang:1.23rc2

# Use apt-get instead of apk for Debian-based images
# Download git to get the dependencies from the go.mod file
RUN apt-get update && apt-get install -y git

# Set the current working directory inside the container
ENV GO111MODULE=on
# Set the GOPATH to /go
ENV GOPATH /go
# Add /go/bin to the PATH
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
# Create the directory structure for the project and set the permissions to 777 for the /go directory and all its subdirectories and files
RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
# Set the working directory to the project directory
WORKDIR $GOPATH/src/backend-github
# Copy the go.mod and go.sum files to the working directory
COPY go.mod .
COPY go.sum .
# Download the dependencies
RUN go mod tidy
# Copy the source code to the working directory
COPY . .
# Build the application
RUN GOOS=linux go build -o /go/bin/app
# Run the application
ENTRYPOINT ["/go/bin/app"]

EXPOSE 3000