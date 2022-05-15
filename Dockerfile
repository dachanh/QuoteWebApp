FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
ENV GO111MODULE off
WORKDIR /app/
ADD ./main /app/
ENTRYPOINT ["./main"]