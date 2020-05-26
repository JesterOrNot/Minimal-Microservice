FROM golang:1.14 AS base

WORKDIR /src/hello-world

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

# Compile the application to /app.
RUN go build -o /app -v ./cmd/hello-world

FROM scratch

COPY --from=base /app /app

CMD [ "/app" ]
