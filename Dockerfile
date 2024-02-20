FROM cgr.dev/chainguard/wolfi-base
RUN apk update

WORKDIR ~/app
ADD dist/ ~/app
CMD ./komeet