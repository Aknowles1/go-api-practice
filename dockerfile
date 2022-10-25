FROM golang:1.19-alpine

RUN apk add --no-cache git

# Set the Current Working Directory inside the container
WORKDIR /app/api

# We want to populate the module cache based on the go.{mod,sum} files.
COPY ./api-code/go.mod .
COPY ./api-code/go.sum .

RUN go mod download

COPY --chmod=777 ./api-code/ .

# Build the Go app
RUN go build -o ./out/api .


# This container exposes port 8080 to the outside world
EXPOSE 8080

# Run the binary program produced by `go install`
CMD ["./out/api"]