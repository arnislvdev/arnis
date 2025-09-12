# GitHub Pages Deployment Guide

## Quick Setup

1. **Push to GitHub:**
   ```bash
   git add .
   git commit -m "Add SEO optimization and GitHub Pages support"
   git push origin main
   ```

2. **Enable GitHub Pages:**
   - Go to your repository settings
   - Navigate to "Pages" section
   - Select "GitHub Actions" as source
   - The workflow will automatically deploy on every push to main

## Custom Domain (Optional)

If you want to use a custom domain:

1. Add your domain to the `cname` field in `.github/workflows/deploy.yml`
2. Create a `CNAME` file in your repository root with your domain
3. Configure DNS settings with your domain provider

## File Structure

```
oldschool-portfolio/
├── .github/workflows/deploy.yml  # GitHub Actions workflow
├── public/icons/                 # Favicon and app icons
├── static/                      # CSS and static assets
├── templates/                   # HTML templates
├── robots.txt                   # SEO robots file
├── sitemap.xml                  # SEO sitemap
├── main.go                      # Go server
└── go.mod                       # Go module
```

## SEO Features Added

- ✅ Comprehensive meta tags
- ✅ Open Graph tags for social sharing
- ✅ Twitter Card support
- ✅ Structured data (JSON-LD)
- ✅ Favicon and app icons
- ✅ Robots.txt and sitemap.xml
- ✅ Canonical URLs
- ✅ Theme colors

## Icons Setup

Replace the placeholder files in `public/icons/` with actual icons:

- `favicon.ico` - Main favicon
- `favicon-16x16.png` - 16x16 favicon
- `favicon-32x32.png` - 32x32 favicon
- `apple-touch-icon.png` - 180x180 Apple touch icon

Generate icons at: https://favicon.io/ or https://realfavicongenerator.net/

## URL Structure

Your site will be available at:
- `https://arnislvdev.github.io/oldschool-portfolio/`
- `https://arnislvdev.github.io/oldschool-portfolio/robots.txt`
- `https://arnislvdev.github.io/oldschool-portfolio/sitemap.xml`
- `https://arnislvdev.github.io/oldschool-portfolio/icons/favicon.ico`
