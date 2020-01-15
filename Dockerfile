FROM alpine:latest
MAINTAINER "lhlyu"
ADD main /app/main
RUN mkdir /app/conf
ADD conf/config.yaml /app/conf/config.yaml
RUN mkdir /app/log
RUN chmod 777 -R /app
ENV LANG en_US.UTF-8
WORKDIR /app
CMD ["./main"]

# 创建容器运行
# docker build -t libra .
# docker run -itd -p 9111:8080 -v /logs/libra-log:/app/log libra