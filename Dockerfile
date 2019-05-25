FROM golang:1.12
WORKDIR /app
COPY ./ ./
RUN GOOS=linux GOARCH=arm go build -o demo-go

FROM arm32v7/alpine
COPY --from=0 /app/demo-go /bin/demo-go
ENTRYPOINT ["/bin/demo-go"]
EXPOSE 8080
