# Stage1 , build image manifest
FROM alpine:3.17.0 as build
LABEL "GoRestAPI"="using alpine base image"
LABEL "Name"="Muhammad Ryan Firmansyah"

RUN apk update && apk add --no-cache git go curl
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o /bin/server

# Multi Stage2, serve result from go build
FROM alpine:3.17.0
WORKDIR /app
COPY --from=build /bin/server /app/binary
COPY --from=build /app/.env .
EXPOSE 8080
ENTRYPOINT ["/app/binary"]
