FROM golang:1.14-alpine AS gobuilder
ARG IMAGE_TAG=undefined # This should be set via the build process
WORKDIR /go/src/github.com/rajch/ttws
COPY . .
COPY web/ttws/www out/www
RUN CGO_ENABLED=0 go build -o out/ttws \
                        -ldflags "-X 'github.com/rajch/ttws/pkg/webserver.version=v${IMAGE_TAG}'" \
                        cmd/ttws/main.go


FROM scratch AS final
COPY --from=gobuilder /go/src/github.com/rajch/ttws/out/ /
ENTRYPOINT [ "/ttws" ]
