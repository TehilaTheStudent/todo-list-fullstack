#build stage
FROM golang:alpine AS builder
RUN apk add --no-cache git
WORKDIR /go/src/app
COPY . .
RUN go get -d -v ./...
RUN go build -o /go/bin/app -v ./cmd/main.go

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ["/app"]
LABEL Name=todolistfullstack Version=0.0.1
EXPOSE 8080
ENV DB_USER=uaqvcqomhgqkvoi5
ENV DB_PASS=VF3McDK3WNeuMSkFmfzu
ENV DB_NAME=bwo8xnxrjgkyqk1fvfxg
ENV DB_HOST=bwo8xnxrjgkyqk1fvfxg-mysql.services.clever-cloud.com
ENV DB_PORT=3306
ENV FRONTEND_URL=https://todo-list-fullstack-frontend.onrender.com
ENV ENV=production
