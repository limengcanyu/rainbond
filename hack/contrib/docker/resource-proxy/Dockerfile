FROM nginx:1.19

VOLUME ["/data/nginx/cache"]
ENV RELEASE_DESC=__RELEASE_DESC__

COPY resource-proxy.conf /etc/nginx/conf.d/
ADD docker-entrypoint.sh /run/docker-entrypoint.sh
RUN chmod +x /run/docker-entrypoint.sh
ENTRYPOINT [ "/run/docker-entrypoint.sh" ]
CMD ["nginx", "-g", "daemon off;"]