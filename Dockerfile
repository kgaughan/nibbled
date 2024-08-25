FROM gcr.io/distroless/static:latest

LABEL org.opencontainers.image.title=nibbled
LABEL org.opencontainers.image.description="A gopher server in Go. Almost meta."
LABEL org.opencontainers.image.vendor="Keith Gaughan"
LABEL org.opencontainers.image.licenses=MIT
LABEL org.opencontainers.image.url=https://github.com/kgaughan/nibbled
LABEL org.opencontainers.image.source=https://github.com/kgaughan/nibbled
LABEL org.opencontainers.image.documentation=https://kgaughan.github.io/nibbled/

# From https://packages.debian.org/sid/media-types
# The file is licensed as public domain.
COPY contrib/mime.types /etc/

COPY nibbled .
EXPOSE 70
ENTRYPOINT ["/nibbled", "--hostname", "0.0.0.0", "--port", "70"]
CMD ["--root", "/srv/gopher"]
