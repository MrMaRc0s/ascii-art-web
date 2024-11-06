# ASCII Art Web

## Description
ASCII Art Web is a web-based application that allows users to convert text into ASCII art using various banner styles. Users can enter text, choose a banner style, and see the generated ASCII art displayed on the webpage. This project builds upon a command-line ASCII art generator by adding an HTTP server and a web GUI.

## Authors
- [Konstantinos Petroutsos, Christos Markos, Sokratis Aggelakopoulos]

## Usage

### Prerequisites
- [Install Go](https://go.dev/dl/) (version 1.18 or later is recommended)

### Running the Server

1. **Clone the Repository** or set up your project directory:
    ```
    ascii-art-web/
    ├── main.go
    ├── templates/
    │   ├── index.html
    │   └── result.html
    ├── banners/
    │    ├── shadow.txt
    │    ├── standard.txt
    │    └── thinkertoy.txt
    ├── ascii_printer.go
    ├── ascii_art_handler.go
    ├── go.mod
    ├── load_banner.go
    ├── main_page.go
    ` ``

2. **Navigate to the Project Directory**:
    ```sh
    cd ascii-art-web
    ```

3. **Run the Server**:
    ```sh
    go run .
    ```

4. **Open a Web Browser** and navigate to:
    ```
    http://localhost:8080
    ```

5. **Use the Web Interface**:
   - Enter your text.
   - Select a banner style.
   - Click "Generate ASCII Art" to display the ASCII art on the page.

6. **Stop the Server** by pressing `Ctrl+C` in the terminal.

## Implementation Details: Algorithm

### Algorithm Overview

1. **Banner Loading**:
   - Based on user input, the program loads a banner file (`standard`, `shadow`, or `thinkertoy`) from the `banners` directory.
   - The selected banner file is read and split into individual lines, which represent ASCII art for each character.

2. **ASCII Art Generation**:
   - The text input is split by literal `"\n"` sequences for multi-line support.
   - Each line is processed recursively with `PrintAsciiArtRecursive`, which:
     - Recursively iterates through each character in the text line.
     - Uses `printSentenceAsciiRecursive` to print each character's ASCII representation by referencing specific line indexes in the banner file.
   - ASCII art lines are built into a result string for display.

3. **Web Server**:
   - The `net/http` package is used to set up HTTP endpoints:
     - `GET /`: Serves the main HTML page for input.
     - `POST /ascii-art`: Handles ASCII art generation requests, rendering the result on a new page.
   - HTML templates (`index.html` and `result.html`) manage user input and result display.

4. **Error Handling**:
   - Appropriate HTTP status codes (e.g., 404, 500) are used for missing files or server errors.
   - Input validation ensures valid banner selection, with errors displayed on the result page.

This project illustrates recursive ASCII art generation and web-based interactivity in Go.
