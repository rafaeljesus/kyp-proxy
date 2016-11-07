FROM scratch
MAINTAINER Rafael Jesus <rafaelljesus86@gmail.com>
ADD kyp-proxy /kyp-proxy
ENV KYP_GATEWAY_HOST=localhost:3004
ENV KYP_PROXY_PORT=8080
ENTRYPOINT ["/kyp-proxy"]
