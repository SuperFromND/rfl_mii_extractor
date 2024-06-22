package main

import (
   "bufio"
   "fmt"
   "io"
   "os"
   "strings"
   "bytes"
   "io/ioutil"
   "regexp"

   "golang.org/x/text/encoding/unicode"
   "golang.org/x/text/transform"
)

// magic constants
const header_length = 4
const mii_data_end_pos = 7424
const mii_data_entry_length = 74

func main() {
    // attempt to load the database file from the local dir
    file, err := os.Open("RFL_DB.dat")
    
    if err != nil {
        fmt.Println(err)
        fmt.Println("Couldn't find/open RFL_DB.dat. Place this file in the same folder that this executable is stored in.\nPress any key to exit (or just close this window).")
        fmt.Scanln()
        return
    }
    
    defer file.Close()
    
    // get the file size
    stat, err := file.Stat()
    if err != nil {
        fmt.Println(err)
        return
    }
    
    // read the file into a byte array
    file_data := make([]byte, stat.Size())
    _, err = bufio.NewReader(file).Read(file_data)
    
    if err != nil && err != io.EOF {
        fmt.Println(err)
        return
    }
    
    // iterate over every mii entry in the database
    for i := header_length; i < mii_data_end_pos - mii_data_entry_length; i += mii_data_entry_length {
        mii_data := file_data[i:i+mii_data_entry_length]
        fmt.Println("Extracting Mii data at:", i)
        
        // extract name bytes
        mii_name_data := mii_data[2:22]
        
        // trashy boilerplate to convert UTF16-BE text to something golang can use
        win16be := unicode.UTF16(unicode.BigEndian, unicode.IgnoreBOM)
        utf16bom := unicode.BOMOverride(win16be.NewDecoder())
        unicodeReader := transform.NewReader(bytes.NewReader(mii_name_data), utf16bom)
        decoded, err := ioutil.ReadAll(unicodeReader)
        regexr :=  regexp.MustCompile(`([^A-Za-z0-9.! ]+)`)
        
        mii_name := regexr.ReplaceAllString(strings.Replace(string(decoded), "\r\n", "\n", -1), "")
        
        if len(strings.TrimSpace(mii_name)) == 0 {
            mii_name = "unknown_mii_" + string(i)
        }
        
        // check to see if data is all-zeroes to prevent blank files from being outputted
        contains_nonzeroes := false
        
        for _, j := range mii_data {
            if j != 0 {
                contains_nonzeroes = true
            }
        }
        
        if !contains_nonzeroes {
            fmt.Println("-- Data is blank. Skipping...")
            continue
        }
        
        // dump bytes to file
        filename := mii_name + ".mii"
        err = os.WriteFile(filename, mii_data, 0666)
        if err != nil {
            fmt.Println(err)
            continue
        }
    }
    
    fmt.Println("\nMii data files should be extracted now. Thank you for using this tool! <3\nPress any key to exit (or just close this window).")
    fmt.Scanln()
    os.Exit(0)
}