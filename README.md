# GOTH Stack 🐈‍⬛🖤

## Overview

This project is a template for the GOTH stack, which combines:

- **Go**: A statically typed, compiled language known for its simplicity and efficiency.
  - **Echo**: A high performance, minimalist Go web framework.
  - **templ**: An HTML templating language for Go that helps in creating reusable UI components.
- **Tailwind CSS**: A utility-first CSS framework for rapidly building custom user interfaces.
- **HTMX**: A lightweight JavaScript library that allows you to access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML.

## Features

- Fast and efficient backend with Go and Echo framework
- Dynamic HTML templating with templ
- Interactive UI updates without full page reloads using HTMX
- Responsive and customizable styling with Tailwind CSS
- Easy-to-use routing and middleware with Echo

## Prerequisites

- Go 1.16 or later
- Node.js and npm
- templ CLI tool

## Setup

1. Clone the repository:

   ```
   git clone https://github.com/yourusername/goth-stack-template.git
   cd goth-stack-template
   ```

2. Install Go dependencies:

   ```
   go mod tidy
   ```

3. Install Node.js dependencies:

   ```
   npm install
   ```

4. Generate templ files:

   ```
   templ generate
   ```

5. Build Tailwind CSS:
   ```
   npm run build-css
   ```

## Running the Application

1. Start the Go server:

   ```
   go run main.go
   ```

2. In a separate terminal, watch for CSS changes:

   ```
   npm run build-css
   ```

3. Open your browser and navigate to `http://localhost:8080`

## Project Structure

- `main.go`: Main entry point of the Go application, contains Echo server setup
- `templates/`: Contains templ files for HTML templating
- `static/`: Static assets including CSS files
- `tailwind.config.js`: Tailwind CSS configuration
- `package.json`: Node.js project configuration and scripts

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
