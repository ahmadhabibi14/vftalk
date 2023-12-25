#!/usr/bin/env bash
set -x

SERVER=root@202.10.36.136

rsync --delete -a \
  --exclude=".git" \
  --exclude=".github" \
  --exclude="node_modules" \
  --exclude="tmp" \
  --exclude="_tmpdb"

rsync -avz --progress . $SERVER:/root/vftalk