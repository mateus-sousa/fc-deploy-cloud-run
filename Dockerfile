FROM golang:1.21 as build
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o cloudrun

FROM scratch
WORKDIR /app
COPY --from=build /app/cloudrun .
EXPOSE 8080
ENV ENVIRONMENT=PROD
ENV WEATHER_TOKEN=b31e34e7d8d842b4849182050243004
ENTRYPOINT [ "./cloudrun" ]
