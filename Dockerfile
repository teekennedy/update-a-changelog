FROM devopsworks/golang-upx:1.23.2 AS builder

ENV CGO_ENABLED=0 \
  GOOS=linux

WORKDIR /src
COPY . .

# Go build options explanation:
# -a: Force rebuilding of packages that are already up-to-date.
# -ldflags: Flags passed to 'go tool link' command.
#   -s: Omit the symbol table and debug information.
#   -w: Omit the DWARF symbol table.
#   -extldflags: Flags to pass to external linker.
#     -static: Do not link to shared libraries.
# -tags netgo: Use pure Golang DNS resolver
# -o /bin/cmd: Output binary as /bin/cmd
RUN go build \
  -a \
  -ldflags "-s -w -extldflags '-static'" \
  -tags netgo \
  -o /bin/cmd \
  . \
  && strip /bin/cmd \
  && upx -q -9 /bin/cmd

RUN echo "nobody:x:65534:65534:Nobody:/:" > /etc_passwd



FROM scratch

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc_passwd /etc/passwd
COPY --from=builder --chown=65534:0 /bin/cmd /cmd

USER nobody
ENTRYPOINT ["/cmd"]
