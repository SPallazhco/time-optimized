FROM golang:1.21 as builder

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o timeslot-optimizer .


FROM debian:bookworm-slim

WORKDIR /app
COPY --from=builder /app/timeslot-optimizer .

EXPOSE 8080

CMD ["./timeslot-optimizer"]