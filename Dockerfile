FROM golang:alpine AS build
WORKDIR /build
ENV CGO_ENABLED=0
COPY . .
RUN go mod tidy
RUN go test -v ./...
RUN GOOS=linux GOARCH=amd64 go build -o food .


FROM alpine:latest
WORKDIR /app
COPY --from=build /build/food .

ENV DATABASE_HOST=localhost
ENV DATABASE_PORT=3306
ENV DATABASE_USER=username
ENV DATABASE_PASS=password
ENV DATABASE_NAME=food

ENV EMAIL_HOST=smtp.example.com
ENV EMAIL_PORT=587
ENV EMAIL_ACCOUNT=sender@example.com
ENV EMAIL_RECIPIENTS=r1@example.com,r2@example.com
ENV EMAIL_PASS=password

CMD ["./food", "-plates", "-plan"]
