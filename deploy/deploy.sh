#!/usr/bin/env bash
set -x

SERVER=root@202.10.36.136

./upload-to-server.sh &&
ssh $SERVER bash /root/vftalk/reload-service.sh
