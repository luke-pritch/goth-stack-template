# GOTH Stack Template

## Overview

This project is a template for the GOTH stack, which combines:

- **Go**: A statically typed, compiled language known for its simplicity and efficiency.
- **templ**: An HTML templating language for Go that helps in creating reusable UI components.
- **HTMX**: A lightweight JavaScript library that allows you to access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML.
- **Tailwind CSS**: A utility-first CSS framework for rapidly building custom user interfaces.

## Features

- Fast and efficient backend with Go
- Dynamic HTML templating with templ
- Interactive UI updates without full page reloads using HTMX
- Responsive and customizable styling with Tailwind CSS

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

- `main.go`: Main entry point of the Go application
- `templates/`: Contains templ files for HTML templating
- `static/`: Static assets including CSS files
- `tailwind.config.js`: Tailwind CSS configuration
- `package.json`: Node.js project configuration and scripts

## Potential Use Cases

1. **Rapid Prototyping**: Quickly build and iterate on web applications with a powerful backend and dynamic frontend.
2. **Single Page Applications (SPAs)**: Create SPAs with smooth, AJAX-powered interactions without heavy JavaScript frameworks.
3. **Content Management Systems (CMS)**: Build custom CMS with Go's performance and HTMX's simplicity.
4. **Real-time Applications**: Leverage Go's concurrency and HTMX's WebSocket capabilities for real-time features.
5. **Microservices Frontend**: Use as a template for building frontends for microservices architectures.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
