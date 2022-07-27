FROM alpine:latest

ARG TURING_USER="turing"
ARG TURING_USER_GROUP="app"
ARG PLUGIN_DIST_PATH="bin/plugin"

RUN addgroup -S ${turing_user_group} \
    && adduser -S ${turing_user} -G ${turing_user_group} -H

ENV PLUGIN_NAME "xp"
ENV PLUGINS_DIR "/app/plugins"

COPY --chown=${TURING_USER}:${TURING_USER_GROUP} ${PLUGIN_DIST_PATH} /go/bin/plugin

CMD ["sh", "-c", "cp /go/bin/plugin ${PLUGINS_DIR}/${PLUGIN_NAME:?variable must be set}"]
