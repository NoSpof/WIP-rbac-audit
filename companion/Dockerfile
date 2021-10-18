FROM alpine:3.12 as builder
ARG K8S_VERSION=v1.19.3
COPY entrypoint.sh /app/
RUN set -x                  && \
    apk --update upgrade    && \
    apk add ca-certificates bash && \
    rm -rf /var/cache/apk/* && \
    wget -O /kubectl https://storage.googleapis.com/kubernetes-release/release/$K8S_VERSION/bin/linux/amd64/kubectl && \
    chmod +x /kubectl && \
    chmod a+x /app/entrypoint.sh
VOLUME [ "/data" ]
CMD [ "/app/entrypoint.sh" ]