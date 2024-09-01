# GOTH Stack Template

## Overview

This project is a template for the GOTH stack, which combines:

- **Go**: A statically typed, compiled language known for its simplicity and efficiency.
- **Echo**: A high performance, minimalist Go web framework.
- **templ**: An HTML templating language for Go that helps in creating reusable UI components.
- **HTMX**: A lightweight JavaScript library that allows you to access AJAX, CSS Transitions, WebSockets and Server Sent Events directly in HTML.
- **Tailwind CSS**: A utility-first CSS framework for rapidly building custom user interfaces.

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

## Potential Use Cases

1. **Rapid Prototyping**: Quickly build and iterate on web applications with Echo's powerful backend and HTMX's dynamic frontend.
2. **Single Page Applications (SPAs)**: Create SPAs with smooth, AJAX-powered interactions without heavy JavaScript frameworks.
3. **RESTful APIs**: Leverage Echo's routing and middleware capabilities to build robust API services.
4. **Content Management Systems (CMS)**: Build custom CMS with Go's performance, Echo's simplicity, and HTMX's interactivity.
5. **Real-time Applications**: Use Echo's WebSocket support alongside HTMX for real-time features.
6. **Microservices Frontend**: Use as a template for building frontends for microservices architectures, with Echo handling backend communication.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is open source and available under the [MIT License](LICENSE).
