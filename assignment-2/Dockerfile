FROM golang:1.20 AS builder

WORKDIR /app
# compile the apps in the builder, still stuck on cross-compile from windows OS to Unix
COPY . .
RUN rm -f go.mod go.sum
RUN go mod init github.com/IlhamEndianto/Hacktiv8/assignment-2
RUN go mod tidy
RUN go build -tags netgo -a -v -o ./binary/assignment-2 .


FROM ubuntu:20.04
WORKDIR /app
COPY --from=builder /app/binary/assignment-2 .
COPY --from=builder /app/template/ template/.
RUN ls
EXPOSE 8080

CMD ["./assignment-2"]