FROM alpine:3.12
RUN apk add --no-cache curl  && rm -rf /var/cache/apk/*
RUN  curl -sSfL https://raw.githubusercontent.com/jckuester/awsls/master/install.sh | sh -s v0.5.1
ENTRYPOINT ["awsls"]
