FROM golang:1.14 AS base

WORKDIR /src/hello-world

# Install dependencies in go.mod and go.sum
COPY go.mod go.sum ./
RUN go mod download

# Copy rest of the application source code
COPY . ./

# Compile the application to /app.
RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' -o /app -v ./src

FROM scratch

COPY --from=base /app /app

EXPOSE 8080

CMD [ "/app" ]
