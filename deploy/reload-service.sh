#!/usr/bin/env bash

cp /root/vftalk/deploy/vftalk.service /lib/systemd/system/ &&
systemctl daemon-reload &&
systemctl enable vftalk &&
systemctl restart vftalk
