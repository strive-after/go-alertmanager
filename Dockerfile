FROM centos:centos7.6.1810
RUN  mkdir /webhook
COPY webhook /webhook/
WORKDIR webhook
CMD /webhook/webhook