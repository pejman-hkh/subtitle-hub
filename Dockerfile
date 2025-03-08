FROM alpine:latest
RUN printf "http://mirror.yandex.ru/mirrors/alpine/v3.17/main\nhttp://mirror.yandex.ru/mirrors/alpine/v3.17/community" > /etc/apk/repositories; \
apk update && apk add --no-cache supervisor

COPY supervisord.conf /etc/supervisord.conf
CMD ["/usr/bin/supervisord", "-c", "/etc/supervisord.conf"]