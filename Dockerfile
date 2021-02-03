ARG GO_VERSION=1.15.7
ARG APP_VENDOR=
ARG REPO_NAME=""
ARG APP_NAME="resourcegroupstaggingapi"
ARG APP_PATH="/go/src/internal/unfor19/aws-sdk-golang-examples/resourcegroupstaggingapi"
# Target executable file:  /app/main 


# Dev
FROM golang:${GO_VERSION}-alpine AS dev
RUN apk add --update git
ARG APP_NAME
ARG APP_PATH
ENV APP_NAME="${APP_NAME}" \
    APP_PATH="${APP_PATH}" \
    GOOS="linux"
WORKDIR "${APP_PATH}"
COPY . "${APP_PATH}"
ENTRYPOINT ["sh"]

# Pass ARGs to next stage
ARG APP_NAME
ARG APP_PATH

# Build
FROM dev as build
ARG APP_NAME
ARG APP_PATH
RUN go mod download
RUN mkdir -p "/app/" && go build -o "/app/main"
ENTRYPOINT [ "sh" ]

# App
FROM alpine AS app
WORKDIR "/app/"
COPY --from=build "/app/main" ./
RUN addgroup -S "appgroup" && adduser -S "appuser" -G "appgroup" && \
    chown "appuser":"appgroup" main
USER "appuser"
ENTRYPOINT ["./main"]
CMD ""
