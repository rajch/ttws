FROM scratch

COPY probestest /
ENTRYPOINT [ "/probestest" ]