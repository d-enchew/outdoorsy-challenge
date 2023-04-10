FROM golang:1.18

WORKDIR /app
COPY . .

RUN go build

# Set environment variables
ENV API_PORT 1212
ENV DB_CONNECTION_URL postgresql://postgres:root@host.docker.internal/postgres?sslmode=disable

RUN chmod 777 ./outdoorsy

# Expose port 1212 for the application
EXPOSE 1212

# Run
CMD ["./outdoorsy"]
