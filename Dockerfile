# https://medium.com/@chemidy/create-the-smallest-and-secured-golang-docker-image-based-on-scratch-4752223b7324
FROM golang as builder

WORKDIR /app

# Cache the go modules
COPY go.mod .
COPY go.sum .

RUN go mod download

# Copy the rest of the files
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o /build

# Multi stage building
FROM scratch
COPY --from=builder /build /app

# COPY ./config/config.yml /config/
EXPOSE 50051
ENTRYPOINT ["./app"]