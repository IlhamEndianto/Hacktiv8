FROM golang:1.18 AS builder

WORKDIR /app
COPY . .

FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/binary/assignment-2 .
COPY --from=builder /app/template/ app/template/.
RUN ls
EXPOSE 8080

CMD ["./assignment-2"]