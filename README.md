## An implementation of Go, Fiber, Handlebars, with WebSocket connection

### Tech stack:
- [Go Fiber](https://gofiber.io)
- [Handlebars](https://handlebarsjs.com/)
- [TailwindCSS](https://tailwindcss.com/)

### How to start development:

```shell
# Install dependencies
go get

bun install

docker-compose up -d

# Create a new docker network
docker network create vftalk-network

# use Air live reload to start web server
air

# or you can do manually
go run main.go

# start nodejs for build javascript or css stuff
bun dev
```

### MariaDB

```shell
docker exec -it vftalk-db mariadb -u root -p
```

### Docker
```shell
# Remove all containers
docker rm -f $(docker ps -aq)

# Remove containers
docker-compose down
```

### TODO:
- [ ] User edit progile
- [ ] Limit online users to only 20
- [x] Login handler
- [x] API rate limit
- [x] Login page
- [x] Fix layout to not overflow
- [x] Responsive to mobile device
- [ ] Turn to Progressive Web App (PWA)
- [x] Add Database to store user data
- [ ] Add Database to store chats
- [x] User Info
- [ ] Add Tenor API for stickers and GIFs
- [ ] Upload picture in chats, and sticker from Tenor
- [x] Use mailhog for SMTP in development, add to docker
- [ ] Deploy to server
- [x] Add MariaDB to docker-compose
- [ ] Use docker-mailserver for SMTP in production
- [x] User can edit profile picture
- [ ] Image editor for update profile picture, specify for 1:1 ratio
- [ ] Direct message
- [ ] Notification
- [ ] Active user list