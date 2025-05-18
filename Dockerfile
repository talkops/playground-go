FROM golang:1.24-alpine AS base
ENV TALKOPS_SOCKET=/tmp/talkops.sock \
    TALKOPS_STDERR=/tmp/talkops.stderr.log \
    TALKOPS_STDOUT=/tmp/talkops.stdout.log
RUN apk add --no-cache nodejs npm && \
    npm install -g pm2@6.0.6 talkops-client@1.0.1 && \
    mkdir /.cache && \
    mkdir /.pm2 && \
    mkdir /app && \
    mkdir /data && \
    chown -R 1000:1000 /.cache /.pm2 /app /data
WORKDIR /app

FROM base AS dev
USER 1000:1000
VOLUME [ "/app" ]
ENTRYPOINT [ "./entrypoint.sh" ]
CMD ["pm2-runtime", "ecosystem.dev.config.cjs"]

FROM base AS prod
COPY ecosystem.prod.config.cjs go.mod go.sum ./
RUN go mod download
COPY src src
USER 1000:1000
CMD ["pm2-runtime", "ecosystem.prod.config.cjs"]
