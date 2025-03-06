# HEX to DEC Converter
Hex2Dec is a simple but powerful hexadecimal (HEX) to decimal (DEC) converter.

[![MIT License](https://img.shields.io/badge/License-MIT-green.svg)](https://choosealicense.com/licenses/mit/)
![Website](https://img.shields.io/website?url=https%3A%2F%2Flukaszwojcik.eu/hex2dec)
![version](https://img.shields.io/badge/version-1.0c-blue)
![Golang](https://img.shields.io/badge/-Golang-00ADD8?logo=Go&logoColor=white&style=flat)
[![LinkedIn](https://img.shields.io/badge/LinkedIn-Connect-blue?style=social&logo=linkedin)](https://www.linkedin.com/in/lukasz-michal-wojcik)

## Contents
- [Application](#application)
- [Why is it worth using?](#using)
- [Installation](#installation)
- [Usage/Examples](#usageexamples)
- [Documentation](#documentation)
- [Bugs](#bugs)
- [Author](#author)
- [Site](#site)
- [Summary](#summary)
- [Download](#download)
- [License](#license)

---

## Application

**HEX to DEC Converter** is a simple yet powerful command-line tool written in Go (Golang) that converts hexadecimal (HEX) values to decimal (DEC) values. It supports both single HEX values and batch processing from a file. The tool is designed for developers, engineers, and anyone working with numerical systems.

### Key Features:
- Convert single HEX values to DEC.
- Process multiple HEX values from a file.
- Save the results to an output file.
- Easy-to-use command-line interface.
- Lightweight and fast.

## Using

### Automation:
The program eliminates the need to manually convert HEX values ​​to DEC, which saves time and reduces the risk of errors.

### Flexibility:
Supports both single values ​​and files with multiple HEX values.

### Simplicity:
It is easy to use thanks to a clear text interface and flag support.

### Security:
Includes error handling for invalid input data.
---

## Installation

### Prerequisites
- Go (Golang) installed on your system. You can download it from [here](https://golang.org/dl/).

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/lukaszwojcikdev/hex2dec_converter.git
   ```
2. Navigate to the project directory:
   ```bash
   cd hex2dec_converter
   ```
3. Build the project:
   ```bash
   go build -o hex2dec_converter
   ```
   or
   ```bash
   go build hex2dec_converter.go
   ```
5. Run the executable:
   ```bash
   ./hex2dec_converter
   ```
---

## Usage/Examples

### Convert a Single HEX Value
To convert a single HEX value to DEC, use the `-dh` flag:
```bash
./hex2dec_converter -dh 583DFD
```
Output:
```
5787133
```

### Convert HEX Values from a File
To convert multiple HEX values from a file, use the `-dhf` flag:
```bash
./hex2dec_converter -dhf input.txt
```
- **Input File Format**: HEX values should be separated by commas (`,`), e.g., `108AB2,1BBBFE,6386DE`.
- **Output**: The results will be saved in a file named `input_dec.txt`.

### Display Help
To display the help menu, use the `-h` flag:
```bash
./hex2dec_converter -h
```

### Display Program Information
To display information about the program, use the `-i` flag:
```bash
./hex2dec_converter -i
```

---
## Documentation
Flags:
-h - Display help.
-i - Display information about the program.
-dh - Convert a single HEX value to DEC.
-dhf - Convert HEX values ​​from a file to DEC and write the result to a new file.

Input file format:
HEX values ​​in the file should be separated by commas (,) without spaces.

Example:
```
108AB2,1BBBFE,6386DE
```

For detailed documentation, visit the official project page:  
[https://lukaszwojcik.eu/hex2dec.html](https://lukaszwojcik.eu/hex2dec.html)

---
## Bugs
If you encounter any bugs or issues, please report them [here](https://github.com/lukaszwojcikdev/hex2dec_converter/issues)

---
## Author
- **Łukasz Wójcik**  
  - Website: [https://lukaszwojcik.eu](https://lukaszwojcik.eu)  
  - LinkedIn: [https://www.linkedin.com/in/lukasz-michal-wojcik](https://www.linkedin.com/in/lukasz-michal-wojcik)  
  - GitHub: [https://github.com/lukaszwojcikdev](https://github.com/lukaszwojcikdev)

---
## Site
Visit the official project site for more information:  
[https://lukaszwojcik.eu/hex2dec.html](https://lukaszwojcik.eu/hex2dec.html)

---
## Summary
**HEX to DEC Converter** is a versatile tool for converting hexadecimal values to decimal values. It is designed to be simple, efficient, and easy to use, making it ideal for developers, engineers, and students working with numerical systems.

---
## Download
You can download the latest release of the project from the [GitHub repository](https://github.com/lukaszwojcikdev/hex2dec_converter/releases).

---
## License
This project is licensed under the **MIT License**. See the [LICENSE](LICENSE) file for details.

---
**Enjoy using HEX to DEC Converter!** 🚀
```
