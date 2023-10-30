## A Chat App

### Tech stack:
- [Go Fiber](https://gofiber.io)
- [Handlebars](https://handlebarsjs.com/)
- [TailwindCSS](https://tailwindcss.com/)

### How to start development:

```shell
# Install dependencies
go get

bun install

# use Air live reload to start web server
air

# or you can do manually
go run main.go

# start nodejs for build javascript or css stuff
bun dev
```


### TODO:
- [ ] User management
- [ ] Limit online users to only 20
- [x] Login handler
- [x] API rate limit
- [x] Login page
- [x] Fix layout to not overflow
- [x] Responsive to mobile device
- [ ] Turn to Progressive Web App (PWA)