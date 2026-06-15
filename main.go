package main

import (
	"html/template"
	"log"
	"net/http"
	"time"
)

type Project struct {
	Title       string
	Description string
	URL         string
	Tech        []string
}

type PortfolioData struct {
	Name     string
	Title    string
	Bio      string
	Projects []Project
	Contact  string
	Year     int
}

func main() {
	http.Handle("/static/", securityHeaders(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))
	http.Handle("/icons/", securityHeaders(http.StripPrefix("/icons/", http.FileServer(http.Dir("public/icons")))))

	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
		http.ServeFile(w, r, "robots.txt")
	})
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
		w.Header().Set("Content-Type", "application/xml")
		http.ServeFile(w, r, "sitemap.xml")
	})

	http.HandleFunc("/", homeHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func setSecurityHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	// No external font sources needed — system fonts only.
	w.Header().Set("Content-Security-Policy",
		"default-src 'self'; "+
			"script-src 'self' 'unsafe-inline'; "+
			"style-src 'self' 'unsafe-inline'; "+
			"img-src 'self' data: https:; "+
			"font-src 'self'; "+
			"connect-src 'self'; "+
			"frame-ancestors 'none';")
}

func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	setSecurityHeaders(w)

	data := PortfolioData{
		Name:    "Arnis",
		Title:   "Full Stack Developer & Design Nerd",
		Bio:     "College student studying Computing & IT. I build things for the web — full stack, end to end. Design-driven by nature, meaning I actually care how it looks, not just how it works. When I'm not coding I'm deep in anime, manga, manhwa, manhua, or a novel. Certified otaku. Seriously though, let's build something.",
		Contact: "hi@arnis.wtf",
		Year:    time.Now().Year(),
		Projects: []Project{
			{
				Title:       "arnis.wtf",
				Description: "This site. A clean portfolio built from scratch in Go — no frameworks, no bloat. Just hand-rolled HTML, CSS, and a server that stays out of the way.",
				URL:         "https://arnis.wtf",
				Tech:        []string{"Go", "HTML", "CSS"},
			},
			{
				Title:       "Coming Soon",
				Description: "A project I'm currently working on. It's a web app that will change the way you think about productivity. Stay tuned for more details!",
				URL:         "#",
				Tech:        []string{"Next.js", "TypeScript", "Tailwind CSS", "Go"},
			},
		},
	}

	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err = tmpl.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
