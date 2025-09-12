# Old School Portfolio

A minimalist, old-school styled portfolio website built with Go and CSS, inspired by the early days of the web.

## Features

- Clean, typography-focused design
- Monospace fonts (Courier New)
- Minimal color palette (black, white, blue)
- Responsive layout
- Terminal-style elements
- ASCII art decorations
- Simple navigation
- Newsletter signup form
- Project showcase
- Contact information

## Project Structure

```
oldschool-portfolio/
├── main.go              # Go web server
├── go.mod              # Go module file
├── templates/
│   └── index.html      # HTML template
├── static/
│   └── css/
│       └── style.css   # Old-school CSS styles
└── README.md           # This file
```

## Getting Started

1. **Install Go** (if not already installed):
   - Download from [golang.org](https://golang.org/dl/)
   - Follow installation instructions for your OS

2. **Clone or download this project**

3. **Navigate to the project directory**:
   ```bash
   cd oldschool-portfolio
   ```

4. **Run the server**:
   ```bash
   go run main.go
   ```

5. **Open your browser** and visit:
   ```
   http://localhost:8080
   ```

## Customization

### Personal Information
Edit the `PortfolioData` struct in `main.go` to customize:
- Your name and title
- Bio description
- Contact email
- Newsletter description
- Projects list

### Styling
Modify `static/css/style.css` to adjust:
- Colors and fonts
- Layout and spacing
- Responsive breakpoints
- Animation effects

### Content
Update `templates/index.html` to add or modify:
- Navigation links
- Sections and content
- ASCII art
- JavaScript functionality

## Design Philosophy

This portfolio embraces the old-school web aesthetic with:
- **Simplicity**: Clean, uncluttered design
- **Typography**: Monospace fonts for that terminal feel
- **Minimalism**: Limited color palette and simple layouts
- **Functionality**: Fast loading and accessible
- **Nostalgia**: Elements reminiscent of early web development

## Browser Support

Works in all modern browsers. Optimized for:
- Chrome/Chromium
- Firefox
- Safari
- Edge

## License

MIT License - feel free to use and modify for your own portfolio.
