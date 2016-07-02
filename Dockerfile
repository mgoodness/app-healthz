FROM scratch
MAINTAINER Michael Goodness <mgoodness@gmail.com>
ADD app-healthz /app-healthz
ENTRYPOINT ["/app-healthz"]
