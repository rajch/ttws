FROM golang:1.14-alpine AS gobuilder
WORKDIR /go/src/github.com/rajch/ttws
COPY . .
RUN CGO_ENABLED=0 go build -o out/ttws cmd/ttws/main.go

FROM scratch AS final
COPY --from=gobuilder /go/src/github.com/rajch/ttws/out/ttws /
ENTRYPOINT [ "/ttws" ]
