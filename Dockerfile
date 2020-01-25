FROM alpine:latest
MAINTAINER "lhlyu"
WORKDIR /app
ADD main ./main
RUN mkdir conf
ADD conf/config.yaml conf/config.yaml
RUN mkdir log
RUN chmod 777 -R /app
ENV LANG en_US.UTF-8
CMD ["./main"]

# 创建容器运行
# docker build -t libra .
# docker run -itd -p 9111:8080 -v /logs/libra-log:/app/log libra
