#!/usr/bin/env bash
# leehom Chen clh021@gmail.com
SCRIPT_PATH=$(realpath "$0")
WEB_PATH="$(dirname "$SCRIPT_PATH")/web"
# echo "$WEB_PATH"
# WEB_PATH="$(dirname "$(dirname "$(dirname "$SCRIPT_PATH")")")"
# CMD="go-bindata -fs -prefix ${WEB_PATH}/web/public/test/ --pkg webtest -o bindata.go ${WEB_PATH}/web/public/test/"
CMD="go-bindata -fs -prefix ${WEB_PATH} --pkg bindataWeb -o bindata.go ${WEB_PATH}"
echo "$CMD"
$CMD
