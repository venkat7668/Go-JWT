From golang:alpine as builder

# Enable go modules
ENV GO111MODULE=on

# Update and install git
RUN apk update && apk add --no-cache git

# Build Delve
#RUN go install github.com/go-delve/delve/cmd/dlv@latest

# Set work dir
WORKDIR /app

# Copy source files to app dir
COPY . .

# Download all the dependencies
RUN go mod download

# Build the project
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -gcflags="all=-N -l" -o ./bin/main .

#TODO: build dlv in builder stage and move to final stage
#TODO: Change final stage image to scratch
FROM golang:1.18

RUN go install github.com/go-delve/delve/cmd/dlv@latest
COPY --from=builder /app/bin/main /

CMD ["dlv", "--listen=:40000", "--headless=true", "--api-version=2", "--accept-multiclient", "exec", "/main"]