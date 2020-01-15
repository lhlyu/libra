FROM alpine:latest
MAINTAINER "lhlyu"
ADD main /app/main
RUN mkdir /app/conf
ADD conf/config.json /app/conf/config.json
RUN chmod 777 -R /app
ENV LANG en_US.UTF-8
CMD ["/app/main"]

# 创建容器运行
# docker build -t libra .
# docker run -itd 8080:8080 libra
