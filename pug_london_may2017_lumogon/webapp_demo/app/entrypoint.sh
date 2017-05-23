#!/bin/bash
sed -i "s/HOSTNAMETOKEN/${HOSTNAME}/g" /usr/share/nginx/html/index.html

nginx -g 'daemon off;'