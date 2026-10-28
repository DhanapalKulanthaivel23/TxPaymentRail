FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
COPY base ./base
COPY controller ./controller
COPY libs ./libs
RUN CGO_ENABLED=0 go build -trimpath -o /api .
FROM scratch
COPY --from=build /api /api
USER 65532:65532
ENV LISTEN_ADDR=:8080
EXPOSE 8080
ENTRYPOINT ["/api"]
