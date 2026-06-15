# arnis.wtf

Personal portfolio — built with Go, no frameworks.

## Run locally

```bash
go run main.go
# → http://localhost:8080
```

## Structure

```
├── main.go              # server + data
├── templates/index.html # HTML template
├── static/css/style.css # styles
└── public/icons/        # favicons
```

## Customize

All personal data (name, bio, projects, contact) lives in the `PortfolioData` struct in `main.go`. Edit there, nothing else needs to change.

## License

MIT