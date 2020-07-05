FROM docker.pkg.github.com/kyoh86/go-check-action/go-check:latest

ENTRYPOINT ["/usr/local/bin/go-check"]
