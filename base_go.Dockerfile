FROM centos:7.6.1810

WORKDIR /go/src/app
# modify CharacterSet TimeZone
RUN yum install kde-l10n-Chinese -y && \
    yum install glibc-common -y && \
    localedef -c -f UTF-8 -i zh_CN zh_CN.utf8 && \
    export LANG=zh_CN.UTF-8 && \
    echo "export LANG=zh_CN.UTF-8" >> /etc/locale.conf && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    yum install wget -y && \
    wget -c https://dl.google.com/go/go1.15.3.linux-amd64.tar.gz && \
    tar zxvf go1.15.3.linux-amd64.tar.gz -C /usr/local/

# 容器环境变量添加，会覆盖默认的变量值
ENV LANG zh_CN.UTF-8
ENV LC_ALL zh_CN.UTF-8
ENV PATH=/go/bin:/usr/local/go/bin:$PATH
ENV GOPATH=/go
ENV GO111MODULE=on
ENV GOPROXY=https://goproxy.cn,direct

CMD ["/bin/bash"]