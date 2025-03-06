package main

import (
    "flag"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
)

func hexToDec(hex string) (string, error) {
    dec, err := strconv.ParseInt(hex, 16, 64)
    if err != nil {
        return "", fmt.Errorf("Invalid HEX value: %s", hex)
    }
    return strconv.FormatInt(dec, 10), nil
}

func hexToDecFile(fileName string) error {
    data, err := os.ReadFile(fileName)
    if err != nil {
        return fmt.Errorf("failed to read file: %w", err)
    }
    hexValues := strings.Split(strings.TrimSpace(string(data)), ",")
    var decValues []string
    for _, hex := range hexValues {
        dec, err := hexToDec(hex)
        if err != nil {
            return fmt.Errorf("conversion error: %w", err)
        }
        decValues = append(decValues, dec)
    }
    outputFileName := strings.TrimSuffix(fileName, ".txt") + "_dec.txt"
    outputData := []byte(strings.Join(decValues, ","))
    err = os.WriteFile(outputFileName, outputData, 0644)
    if err != nil {
        return fmt.Errorf("failed to write file: %w", err)
    }
    return nil
}

func printHelp() {
    fmt.Println("HEX to DEC Converter v1.0c")
    fmt.Println("░█░█░█▀▀░█░█░▀▀▄░█▀▄░█▀▀░█▀▀")
    fmt.Println("░█▀█░█▀▀░▄▀▄░▄▀░░█░█░█▀▀░█░░")
    fmt.Println("░▀░▀░▀▀▀░▀░▀░▀▀▀░▀▀░░▀▀▀░▀▀▀")
    fmt.Println("")
    fmt.Println("(c) by Lukasz Wojcik 2024, 2025")
    fmt.Println("Use: hex2dec [flag] [argument]")
    fmt.Println("")
    fmt.Println("Flags:")
    fmt.Println(" -h   Display help")
    fmt.Println(" -i   Displays program information")
    fmt.Println(" -dh  Converts a HEX value to DEC ")
    fmt.Println(" -dhf Converts HEX values from file and writes DEC to file")
    fmt.Println("")
    fmt.Println("Examples of use:")
    fmt.Println(" hex2dec -dh 583DFD")
    fmt.Println(" hex2dec -dhf input.txt")
    fmt.Println("")
    fmt.Println("HEX values in the file should be separated by a comma (,) with no 'space'.")
    fmt.Println("Example of the HEX system arrangement in the 'inside.txt' file:")
    fmt.Println("108AB2,1BBBFE,6386DE,...,...")
}

func printProgramInfo() {
    fmt.Println("Hex2Dec - HEX to DEC Converter")
    fmt.Println("Author: Lukasz Wojcik")
    fmt.Println("Version: 1.0c")
    fmt.Println("Year: 2025")
    fmt.Println("Site: https://lukaszwojcik.eu/hex2dec.html")
    fmt.Println("Source Code: https://github.com/lukaszwojcikdev/hex2dec_converter.git")
}

func main() {
    convertHex := flag.String("dh", "", "Converts HEX value to DEC")
    help := flag.Bool("h", false, "Displays help")
    info := flag.Bool("i", false, "Displays information about the program")
    convertHexFile := flag.String("dhf", "", "Converts HEX values from a file and writes DEC to the file")
    flag.Parse()

    if *help {
        printHelp()
        return
    }

    if *info {
        printProgramInfo()
        return
    }

    if *convertHex != "" {
        dec, err := hexToDec(*convertHex)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println(dec)
        return
    }

    if *convertHexFile != "" {
        err := hexToDecFile(*convertHexFile)
        if err != nil {
            log.Fatal(err)
        }
        fmt.Println("Conversion completed successfully. Results saved to file: input_dec.txt")
        return
    }
    printHelp()
}