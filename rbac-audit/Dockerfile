FROM golang:alpine3.12 as builder
COPY /src /app
WORKDIR /app
RUN apk add jq && go build
CMD ["./k8s-rbac-audit"]

FROM gcr.io/distroless/base
COPY --from=builder /app/k8s-rbac-audit /
CMD ["/k8s-rbac-audit"]