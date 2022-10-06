From golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# Update and install git
RUN apk update && apk add --no-cache git

# Set work dir
WORKDIR /app

# Copy mod and sum files
COPY go.mod ./
COPY go.sum ./

# Download all the dependencies
RUN go mod download

# Copy source files to app dir
COPY . .

# Build the project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main .

From scratch

COPY --from=builder /app/bin/main .

CMD ["./main"]