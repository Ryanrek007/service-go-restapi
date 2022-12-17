FROM alpine:3.17.0
LABEL "GoRestAPI"="using alpine base image"
LABEL "Name"="Muhammad Ryan Firmansyah"
LABEL "Technical-Test"="Site Reliability Engineering at Bibit.id"

RUN apk update && apk add --no-cache git go curl
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o binary
EXPOSE 8080
ENTRYPOINT ["/app/binary"]
