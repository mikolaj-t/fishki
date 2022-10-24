FROM golang:1.19.0-alpine3.16 as build
RUN mkdir /fishkiServer
WORKDIR /fishkiServer
COPY . .
RUN ls -R
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o main ./cmd

FROM alpine:3.16.0
COPY --from=build /fishkiServer/main /fishkiServer/main
EXPOSE 8080
CMD ["/fishkiServer/main"]