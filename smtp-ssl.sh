docker run --rm -it \
  -v "${PWD}/_docker-data/certbot/certs/:/etc/letsencrypt/" \
  -v "${PWD}/_docker-data/certbot/logs/:/var/log/letsencrypt/" \
  -p 80:80 \
  certbot/certbot certonly --standalone -d mail.vftalk.my.id