#!/bin/sh

echo "[`date`] Running entrypoint script"

cp ./${CONFIG_FILE} ./config.yaml
find config* | grep -v config.yaml | xargs rm -f

echo "[`date`] Start Service '${SERVICE_NAME}' "

./courier-management start ${SERVICE_NAME}
