#!/usr/bin/env bash
set -x

sudo apt install nginx
echo '
map $http_upgrade $connection_upgrade {
  default upgrade;
  ''      close;
}

server {
  listen 80;
  listen [::]:80;

  server_name vftalk.my.id;

  location / {
    proxy_pass http://localhost:8000;
    proxy_http_version 1.1;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
  }

  location /api/room {
    proxy_pass http://localhost:8000/api/room;
    proxy_http_version 1.1;
    proxy_set_header Upgrade $http_upgrade;
    proxy_set_header Connection $connection_upgrade;
    proxy_set_header Host $host;
    proxy_set_header X-Real-IP $remote_addr;
    proxy_set_header X-Forwarded-For $remote_addr;
  }
}

' | sudo tee /etc/nginx/sites-available/vftalk
sudo ln -s /etc/nginx/sites-available/vftalk /etc/nginx/sites-enabled/
sudo nginx -t
sudo systemctl restart nginx
