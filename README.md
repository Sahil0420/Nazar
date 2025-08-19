# Go_movie

```bash
tree
.
├── db
│   └── db.go
├── go.mod
├── go.sum
├── handlers
│   ├── add_movie_handlers.go
│   ├── admin_login_handlers.go
│   ├── dashboard_handler.go
│   └── handlers.go
├── main.go
├── README.md
├── static
│   ├── assets
│   │   ├── broadcast-svgrepo-com.svg
│   │   └── menu-burger-horizontal-svgrepo-com.svg
│   ├── header.css
│   ├── script.js
│   └── style.css
└── templates
    ├── admin_login.html
    ├── base.html
    ├── components
    │   └── header.html
    ├── dashboard.html
    ├── index.html
    └── movie.html
```

**Gomovies** is a Movies Reviews Site Built by using Go lang and using built in support html/template for making dynamic html pages. This project have a modular structure . Postgres DB is used as Database. This will have few pages like index.html for landing page, admin_login.html which will secured by the password. handlers are used for backend.