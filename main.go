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
	// Serve static files (CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	// Serve public assets (icons, images, etc.)
	http.Handle("/icons/", http.StripPrefix("/icons/", http.FileServer(http.Dir("public/icons"))))
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))

	// Serve SEO files
	http.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "robots.txt")
	})
	http.HandleFunc("/sitemap.xml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/xml")
		http.ServeFile(w, r, "sitemap.xml")
	})

	// Handle main page
	http.HandleFunc("/", homeHandler)

	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
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
