#!/bin/bash
TARGET_CONTAINER=webappdemo_app2_1
echo "Simulating hack of ${TARGET_CONTAINER}"
echo " - updating index.html"
docker exec ${TARGET_CONTAINER} cp /app/templates/hacked.html /app/templates/index.html
echo " - adding wanacrypt.exe (fake)"
docker exec ${TARGET_CONTAINER} cp /app/templates/hacked.html /app/templates/wanacrypt.exe