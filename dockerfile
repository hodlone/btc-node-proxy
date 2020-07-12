FROM golang:1.13.12-alpine3.12

RUN apk add --no-cache make gcc g++ zeromq zeromq-dev

# Add Maintainer Info
LABEL maintainer="Ricardo Alonzo<alonzoricardo6@gmail.com>"

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum /app/

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN go build -o main .

# Expose port 4000 to the outside world
EXPOSE 4000

# Command to run the executable
CMD ["/app/main"]