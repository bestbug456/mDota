FROM alpine:latest

MAINTAINER Kenny <bongikairu@gmail.com>

WORKDIR "/opt"

ADD .docker_build/mDota /opt/bin/mDota
ADD ./templates /opt/templates
ADD ./data /opt/data
# ADD ./static /opt/static

CMD ["/opt/bin/mDota"]

