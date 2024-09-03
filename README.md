# GoTTH

## 📌 Overview
This repo is a template for GoTTH stack for developing web applications. The stack contains:
- Go: Golang handles routes using Echo, render Templ's template
- Templ: Creating templates, components such as buttons...
- Tailwind CSS: Styling
- Htmx: It is what it is...

## 🔨 Installation

### Prerequisites

- [Go](https://go.dev/) version >= 1.22.5
- [mingw32-make](https://www.msys2.org/) or make (if you use make, please fix `mingw32-make` to `make` in `.air.toml` and `Makefile`)

### Running
To install dependencies, please type the following command in the terminal:

```bash
go mod tidy
```

After that, you need to have the standalone CLI for Tailwindcss:

- Either install [here](https://tailwindcss.com/blog/standalone-cli), create a new directory called `build` and move the executable inside the `build directory`
- Or if you are on Windows (sucks btw), just type:


```bash
make setup-win
```

This will use curl and do the aforementioned steps automatically

Finnaly, you can run:

```bash
make start
```

This will run Air, a dependency that will check for file changes and reload the webpage, although not as fast as `Live Server`. 

## 📫 Feedbacks
If you have any problems, feel free to open an issue. Thank you for your time...
