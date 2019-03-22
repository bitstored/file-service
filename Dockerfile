FROM golang:alpine as source
WORKDIR /home/server
COPY . .
WORKDIR cmd/file-service
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -mod vendor -o file-service

FROM alpine as runner
LABEL name="bitstored/file-service"
RUN apk --update add ca-certificates
COPY --from=source /home/server/cmd/file-service/file-service /bin/file-service
EXPOSE 8080
ENTRYPOINT [ "file-service" ]