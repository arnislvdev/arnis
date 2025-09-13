#!/bin/bash

# Build the Go application
echo "Building Go application..."
go build -o portfolio main.go

# Start the server in background
echo "Starting server..."
./portfolio &
SERVER_PID=$!

# Wait for server to start
sleep 2

# Create dist directory
mkdir -p dist

# Generate static HTML
echo "Generating static HTML..."
curl -s http://localhost:8080/ > dist/index.html

# Copy static assets
echo "Copying static assets..."
cp -r static dist/
cp -r public dist/
cp _headers dist/
cp _redirects dist/
cp robots.txt dist/
cp sitemap.xml dist/

# Kill the server
echo "Stopping server..."
kill $SERVER_PID

echo "Static site generated in dist/ directory"
echo "You can now deploy the contents of dist/ to any static hosting service"
