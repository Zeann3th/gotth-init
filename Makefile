start:
	air

setup-win:
	curl -sLO https://github.com/tailwindlabs/tailwindcss/releases/download/v3.4.10/tailwindcss-windows-x64.exe
	mkdir build
	mv .\tailwindcss-windows-x64.exe .\build\

tailwind-build:
	.\build\tailwindcss-windows-x64.exe -i internal/css/input.css -o internal/css/output.css --minify


