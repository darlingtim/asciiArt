# ASCII Art Generator

This is a command-line application written in Go that takes a text input and generates graphic ASCII art using customizable banner files.

## Features

- **Text to ASCII Art:** Converts standard text strings into ASCII art banners.
- **Custom Banners:** Supports loading custom `.txt` banner files (defaults to `standard.txt`).
- **Newline Handling:** Parses literal `\n` in the input string to generate multiline ASCII art.
- **Validation:** Automatically validates the integrity of the banner file (checks for exactly 855 lines and consistent character widths).
- **Fallback Character:** Unsupported characters are gracefully replaced with a `#` character.

## Usage

Run the program using the `go run .` command from the root of the project directory.

### Syntax

```bash
go run . "<text_input>" [banner_name]
```

### Examples

**1. Basic Usage (uses default `standard.txt`):**
```bash
go run . "Hello"
```
```text
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
```                               

**2. Using a specific banner:**
```bash
go run . "Hello World" shadow
```

```text
                                                                                    
_|    _|          _| _|                _|          _|                   _|       _|  
_|    _|   _|_|   _| _|   _|_|         _|          _|   _|_|   _|  _|_| _|   _|_|_|  
_|_|_|_| _|_|_|_| _| _| _|    _|       _|    _|    _| _|    _| _|_|     _| _|    _|  
_|    _| _|       _| _| _|    _|         _|  _|  _|   _|    _| _|       _| _|    _|  
_|    _|   _|_|_| _| _|   _|_|             _|  _|       _|_|   _|       _|   _|_|_|  
                                                                                    
```

*(Note: The `.txt` extension is automatically appended if omitted.)*

**3. Handling Newlines:**
```bash
go run . "Hello\nWorld" standard
```
```text
 _    _          _   _          
| |  | |        | | | |         
| |__| |   ___  | | | |   ___   
|  __  |  / _ \ | | | |  / _ \  
| |  | | |  __/ | | | | | (_) | 
|_|  |_|  \___| |_| |_|  \___/  
                                
                                
__          __                 _       _  
\ \        / /                | |     | | 
 \ \  /\  / /    ___    _ __  | |   __| | 
  \ \/  \/ /    / _ \  | '__| | |  / _` | 
   \  /\  /    | (_) | | |    | | | (_| | 
    \/  \/      \___/  |_|    |_|  \__,_| 
                                          
```                                          

## Project Structure

- `main.go`: Entry point, parses command-line arguments, and orchestrates the generation.
- `generateAscii.go`: Handles the mapping of input characters to their ASCII art representation.
- `loadBanner.go`: Reads the banner file and maps it to a Go map (character rune -> ASCII string array).
- `validate.go`: Ensures the provided banner file is structurally sound before processing.
- `splitlines.go`: Helper function to consistently split the banner file data by newlines.

## Author

- **Timothy Ododo** - [follow me on github](https://github.com/darlingtim)