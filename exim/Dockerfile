FROM docker.io/alpine:latest

RUN apk --no-cache add exim tini && \
    mkdir /var/spool/exim && \
    chmod 777 /var/spool/exim && \
    ln -sf /dev/stdout /var/log/exim/mainlog && \
    ln -sf /dev/stderr /var/log/exim/panic && \
    ln -sf /dev/stderr /var/log/exim/reject && \
    chmod 0755 /usr/sbin/exim

COPY exim.conf /etc/exim/exim.conf
RUN chmod 664 /etc/exim/exim.conf

USER exim

ENV LOCAL_DOMAINS=@ \
    RELAY_FROM_HOSTS=10.0.0.0/8:172.16.0.0/12:192.168.0.0/16\
    RELAY_TO_DOMAINS=* \
    RELAY_TO_USERS=""\
    DISABLE_SENDER_VERIFICATION="" \
    HOSTNAME="" \
    SMARTHOST="" \
    SMTP_PASSWORD="" \
    SMTP_USERDOMAIN="" \
    SMTP_USERNAME=""


#SSL_CERT_PATH : SSL_KEY_PATH

ENTRYPOINT ["/sbin/tini", "--"]
CMD ["exim", "-bdf", "-q15m"]
