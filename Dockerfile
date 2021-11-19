FROM golang:1.17.1 AS build
WORKDIR /go_app
COPY app.go ./
RUN CGO_ENABLED=0 go build app.go

FROM alpine:3.14.3
WORKDIR /academy
COPY --from=build /go_app/app ./
EXPOSE 8080
ENTRYPOINT [ "./app" ] 
RUN rm -rf /bin

