FROM scratch

COPY ttws /
ENTRYPOINT [ "/ttws" ]