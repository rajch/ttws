FROM scratch

COPY ldgen /
ENTRYPOINT [ "/ldgen" ]