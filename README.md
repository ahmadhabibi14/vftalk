# VFtalk - Chat App

## Tech stack:
- Programming Language: [Go-Lang](https://go.dev), CSS, [JavaScript](https://www.javascript.com/), SQL, Bash
- Http Router: [Go Fiber](https://gofiber.io)
- Template Engine: [Handlebars](https://handlebarsjs.com/)
- UI Library: [TailwindCSS](https://tailwindcss.com/), [Iconsax](https://iconsax.io/)
- DBMS: [MariaDB](https://mariadb.org/)
- Container: [Docker](https://www.docker.com/)
- CI/CD: Github Action
- Web Server: [NGINX](https://www.nginx.com/)
- SMTP: [Docker-mailserver](https://github.com/docker-mailserver/docker-mailserver), [Mailhog](https://github.com/mailhog/MailHog)

### How to start development:

```shell
##### Install dependencies
go mod tidy
pnpm install

##### Set up docker
docker-compose up -d
docker network create vftalk-network

##### Start App
go run main.go web

##### Start nodejs for build javascript or css stuff
pnpm dev
```

### MariaDB

```shell
##### Login to MariaDB CLI
docker exec -it vftalk-db mariadb -u root -p

##### Database migration
### Install golang-migrate
go install -tags "postgres,mysql" github.com/golang-migrate/migrate/v4/cmd/migrate@latest

### Create Migration
migrate create -ext sql -dir database/migration migration_state

### migrate up
make migrate-up
## or
go run main.go migrate

### migrate down
make migrate-down
```

### Docker
```shell
# Remove all containers
docker rm -f $(docker ps -aq)

# Remove containers
docker-compose down
```

### Deploy
```shell
cd deploy
##### Execute deploy script, it will automatically do their magic
./deploy.sh

#### or
## Configure github action script for CI/CD
## add github secrets, make sure which variable to store as defined at .github/workflows/deploy.yaml
```

### TODO:
- [x] User edit progile
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
- [ ] Upload picture in chats, and stickers from Tenor API
- [x] Use mailhog for SMTP in development, add to docker
- [x] Deploy to server (must use linux)
- [x] Add MariaDB to docker-compose
- [ ] Use docker-mailserver for SMTP in production
- [x] User can edit profile picture
- [ ] Image editor for update profile picture, specify for 1:1 ratio
- [ ] Compress image after user uploaded it
- [ ] Direct message
- [ ] Notification (server sent event)
- [ ] Active user list (for direct message)
- [x] Database migration
