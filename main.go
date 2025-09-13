package main

import (
	"html/template"
	"log"
	"net/http"
)

type Project struct {
	Title       string
	Description string
	URL         string
	Tech        []string
}

type PortfolioData struct {
	Name       string
	Title      string
	Bio        string
	Projects   []Project
	Contact    string
	Newsletter string
}

func main() {
	// Serve static files (CSS, JS, etc.) with security headers
	http.Handle("/static/", securityHeaders(http.StripPrefix("/static/", http.FileServer(http.Dir("static")))))

	// Serve public assets (icons, images, etc.) with security headers
	http.Handle("/icons/", securityHeaders(http.StripPrefix("/icons/", http.FileServer(http.Dir("public/icons")))))
	http.Handle("/public/", securityHeaders(http.StripPrefix("/public/", http.FileServer(http.Dir("public")))))

	// Serve SEO files with security headers
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
		http.ServeFile(w, r, "robots.txt")
	})
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
		w.Header().Set("Content-Type", "application/xml")
		http.ServeFile(w, r, "sitemap.xml")
	})

	// Handle main page
	http.HandleFunc("/", homeHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// setSecurityHeaders sets security headers for HTTP responses
func setSecurityHeaders(w http.ResponseWriter) {
	w.Header().Set("X-Frame-Options", "DENY")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.Header().Set("X-XSS-Protection", "1; mode=block")
	w.Header().Set("Referrer-Policy", "strict-origin-when-cross-origin")
	w.Header().Set("Permissions-Policy", "camera=(), microphone=(), geolocation=()")
	w.Header().Set("Strict-Transport-Security", "max-age=31536000; includeSubDomains; preload")
	w.Header().Set("Content-Security-Policy", "default-src 'self'; script-src 'self' 'unsafe-inline'; style-src 'self' 'unsafe-inline'; img-src 'self' data: https:; font-src 'self' data:; connect-src 'self'; frame-ancestors 'none';")
}

// securityHeaders middleware adds security headers to responses
func securityHeaders(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		setSecurityHeaders(w)
		next.ServeHTTP(w, r)
	})
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	// Set security headers
	setSecurityHeaders(w)

	// Portfolio data
	data := PortfolioData{
		Name:       "Arnis [REDACTED]",
		Title:      "Aspiring Full Stack Developer",
		Bio:        "18-year-old college student passionate about full-stack development and cybersecurity. Learning to build modern web applications while exploring the security aspects of software development.",
		Contact:    "hi@arnis.wtf",
		Newsletter: "I write about web development, design, and the evolution of the internet.",
		Projects: []Project{
			{
				Title:       "My Website Portfolio",
				Description: "A minimalist portfolio website built with Go and CSS, embracing the early web aesthetic.",
				URL:         "https://arnis.wtf",
				Tech:        []string{"Go", "HTML", "CSS"},
			},
			{
				Title:       "Tip Me Bro",
				Description: "Terminal styled web app for tipping me bro",
				URL:         "https://tipmebro.wtf",
				Tech:        []string{"JavaScript", "TypeScript", "Next.js"},
			},
		},
	}

	// Parse and execute template
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
