# বিল্ড স্টেজ
FROM golang:1.21-alpine AS builder

WORKDIR /app

# প্রয়োজনীয় প্যাকেজ
RUN apk add --no-cache git

# go.mod এবং go.sum কপি
COPY go.mod go.sum ./
RUN go mod download

# সব সোর্স কপি
COPY . .

# 🔥 বাইনারি বিল্ড
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o blooddnr ./main.go

# রান স্টেজ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

# 🔥 বাইনারি কপি
COPY --from=builder /app/blooddnr .

# পোর্ট এক্সপোজ
EXPOSE 8080

# 🔥 রান কমান্ড
CMD ["./blooddnr"]