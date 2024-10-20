package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

type File struct {
    bytes int
    lines int
    words int
    characters int
}

func (f *File) parse(f_io *os.File) error {
    reader := bufio.NewReader(f_io)
    for {
        str, err := reader.ReadString('\n')
        if err != nil {break}
        f.bytes += len(str)
        f.lines++
        f.characters += utf8.RuneCountInString(str)
        str = strings.TrimSpace(str)
        str = strings.ReplaceAll(str, "\t", " ")
        words := strings.Split(str, " ")
        for _, w := range words {
            if len(w) != 0 {
                f.words++
            }
        }

    }

    return nil
}

func main() {
    byteCountFlag := flag.Bool("bytes", false, "setting this to use count bytes")
    flag.BoolVar(byteCountFlag, "c", false, "setting this to count bytes")
    lineCountFlag := flag.Bool("lines", false, "setting this to use count lines")
    flag.BoolVar(lineCountFlag, "l", false, "setting this to count lines")
    wordCountFlag := flag.Bool("words", false, "setting this to use count words")
    flag.BoolVar(wordCountFlag, "w", false, "setting this to count words")
    charCountFlag := flag.Bool("characters", false, "setting this to use count characters")
    flag.BoolVar(charCountFlag, "m", false, "setting this to count characters")

    flag.Parse()
    rest := flag.Args()

    var f_io *os.File
    var fileName string
    if len(rest) == 0 {
        f_io = os.NewFile(os.Stdin.Fd(), "standard input")
    } else {
        fileName = rest[0]
        f_file_io, err := os.Open(fileName)
        if err != nil {
            fmt.Println("Error reading file")
            os.Exit(1)
        }
        f_io = f_file_io
    }

    file := File{}
    err := file.parse(f_io)
    if err != nil {
        fmt.Println("error parsing the file: ", err)
        os.Exit(1)
    }
    switch {
        case *byteCountFlag:
            fmt.Printf("%d %s\n", file.bytes, fileName)
        case *lineCountFlag:
            fmt.Printf("%d %s\n", file.lines, fileName)
        case *wordCountFlag:
            fmt.Printf("%d %s\n", file.words, fileName)
        case *charCountFlag:
            fmt.Printf("%d %s\n", file.characters, fileName)
        default:
            fmt.Printf("%d %d %d %s\n", file.lines, file.words, file.bytes, fileName)
    }
}
