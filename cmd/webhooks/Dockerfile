FROM golang:1.18.1-bullseye

ENV JX_HOME /home/jx
RUN apt-get install -y ca-certificates git && \
    mkdir $JX_HOME
# && adduser -disabled-password -u 1000 jx
#USER 1000

COPY entrypoint.sh /home/jx/entrypoint.sh
COPY ./main /home/jx/webhooks
ENTRYPOINT ["/home/jx/entrypoint.sh"]
