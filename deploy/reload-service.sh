#!/usr/bin/env bash
set -x

cp /root/vftalk/deploy/vftalk.service /lib/systemd/system/ &&
systemctl daemon-reload &&
systemctl enable vftalk &&
systemctl restart vftalk &&
systemctl restart nginx