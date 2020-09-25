FROM golang:alpine

# Set necessary environmet variables needed for our image
ENV GOOS=linux \
    GOARCH=amd64\
    BINARY_NAME="redact-api"

# Move to working directory /build
WORKDIR $GOPATH/src/github.com/gufranmirza/redact-api-golang/src

# Copy the code into the container
COPY . .

# Download dependency using go mod
RUN go mod download

# Build the application
RUN go build -o ../bin/${BINARY_NAME} .

# Move to /dist directory as the place for resulting binary folder
WORKDIR /dist

# Copy binary from build to main folder
RUN cp $GOPATH/src/github.com/gufranmirza/redact-api-golang/bin/${BINARY_NAME} .
RUN cp $GOPATH/src/github.com/gufranmirza/redact-api-golang/src/app-config.json .

# Export necessary port
EXPOSE 8001

# Command to run when starting the container
CMD ["/dist/redact-api"]

