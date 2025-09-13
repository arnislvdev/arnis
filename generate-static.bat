@echo off

REM Build the Go application
echo Building Go application...
go build -o portfolio.exe main.go

REM Start the server in background
echo Starting server...
start /B portfolio.exe

REM Wait for server to start
timeout /t 3 /nobreak > nul

REM Create dist directory
if not exist dist mkdir dist

REM Generate static HTML
echo Generating static HTML...
curl -s http://localhost:8080/ > dist/index.html

REM Copy static assets
echo Copying static assets...
xcopy static dist\static\ /E /I /Y
xcopy public dist\public\ /E /I /Y
copy _headers dist\
copy _redirects dist\
copy robots.txt dist\
copy sitemap.xml dist\

REM Kill the server
echo Stopping server...
taskkill /F /IM portfolio.exe > nul 2>&1

echo Static site generated in dist/ directory
echo You can now deploy the contents of dist/ to any static hosting service
pause
