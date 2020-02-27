FROM alpine:latest
MAINTAINER "lhlyu"
WORKDIR /app
RUN mkdir conf && mkdir log && chmod 777 -R /app
COPY main ./main
COPY conf/config.yaml conf/config.yaml
CMD ["./main"]

# 创建容器运行
# docker build -t libra .
# docker run -itd -p 9111:8080 -v /logs/libra-log:/app/log libra
