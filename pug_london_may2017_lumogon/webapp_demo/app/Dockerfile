FROM nginx

LABEL version "1.0"
LABEL description "This is an example webapp"
LABEL org.label-schema.schema-version "1.0"
LABEL org.label-schema.build-date "2016-04-12T23:20:50.52Z"
LABEL org.label-schema.name "webapp_demo-app"
LABEL org.label-schema.vcs-ref "279FA63"
LABEL org.label-schema.vcs-url "https://github.com/johnmccabe/talks/pug_london_may2017_lumogon/webapp_demo/app"
LABEL org.label-schema.vendor "John McCabe"

COPY cover.css /usr/share/nginx/html/
COPY favicon.ico /usr/share/nginx/html/
COPY index.html /usr/share/nginx/html/
COPY dummyhack.html /usr/share/nginx/html/

COPY entrypoint.sh /
ENTRYPOINT ["/entrypoint.sh"]