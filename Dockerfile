FROM alpine:latest

# app name
ENV AppName=lottery-period


# mkdir directory
RUN mkdir -p /app/logs

# copy file to image
COPY ${AppName} /app/
COPY conf     /app/conf
COPY yaml   /app/yaml


# add timezone
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
        && apk --no-cache add ca-certificates tzdata\
        && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
        && echo "Asia/Shanghai" > /etc/timezone \
        && apk del tzdata

WORKDIR /app

# 日志挂载
VOLUME /app/logs

# port
EXPOSE 10002

# label
LABEL maintainer = "gkzy"

CMD ["sh","-c","./$AppName"]