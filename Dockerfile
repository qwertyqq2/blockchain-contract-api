FROM docker.io/golang:1.21-alpine as build
WORKDIR /app
COPY go.mod go.sum /app/
RUN go mod graph | awk '{if ($1 !~ "@") print $2}' | xargs go get
COPY . .
RUN go build -o bin main.go

FROM scratch
COPY conf.yaml ./
COPY --from=build /app/bin /app/bin
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

CMD ["/app/bin"]
