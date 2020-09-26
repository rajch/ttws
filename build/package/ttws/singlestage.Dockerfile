FROM scratch

COPY ttws /
COPY www/ /www
ENTRYPOINT [ "/ttws" ]