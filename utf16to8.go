// Copyright 2018 Awakecard
// Use of this sourcecode is granted under the MIT Licence
// please see the licence file un this directory


//utf16toutf8 is a tool which is designed to convert text based filed from utf-16 little-endian to utf-8
//where there are plans in the future to also allow for big-endian convertion this has not been coverted.
//
// utf16toutf8 has the following package variables:
//
// Infile,(string) this is the file you want to read from
// Outfile,(string) this is the file you want to create
// Errors,([]string) this is a string list, which will keep track of errors
// RunCode,(int) this is a indication of the last error, if there is no error then the run code will remain at zero
//
// External packaged and dependancies
//
// - os
//
// Sample usage
// import {
// 	....
// "utf16to8"
// ....
// }
//
// func xxxx ( xxx ) xxx {
//  ....
//
// utf16to8.Infile = "....."
// utf16to8.Outfile = "...."
// if utf16to8.Convert() {
//		// it worked
//		....
// } else {
//		// it failed
//		....
// }


package utf16to8


import (
	"fmt"
	"io/ioutil"
	"os"
)

var Infile = ""
var Outfile = ""

var Errors = []string{}

var RunCode = 0


///
/// Convert
//  :param void
//  :return bool
//   this is the main function,
//   it will test to see if the Infile and the out file exists, before exiting and then continuing on with the program
//   it will try to write to to a new outfile, if this fails then it will exit
//   it will try to read from the in file if it fails then it will exit
//   it will then try and determine the infiles bom, if the BOM is not uft-16 little endian then it will fail
//   during the conversion process, if there is a character which can not be mapped to utf-8 then the character will be
//   replaced with the hex code
///
func Convert() bool {
	if !fileExits(Infile){
		Errors = append(Errors, "Infile does not Exist please set it and try again, code 1")
		RunCode = 1
	}

	if fileExits(Outfile){
		Errors = append(Errors, "Outfile exits please remove out file and retry, code 2")
		RunCode = 2
	}


	if RunCode > 0 {
		return false
	}




	inFileBytes ,e := ioutil.ReadFile(Infile)
	outFileHandle, _ := os.Create(Outfile)

	if e != nil {
		fmt.Println(e)
		RunCode = 4
		Errors = append(Errors, "Failed to Read Infile, code 4")
		panic (e)
	}


	var line = []byte{}


	for i,c := range( inFileBytes ) {

		if (i%2 == 0){
			if c >= 20 && c <= 126 {
				line = append(line, c)
			}else {
				if c == 10 {
					line = append(line, c)
					outFileHandle.Write(line)
					line = []byte{}
				}else if c == 13{

				}else {
					fmt.Println(c)
				}
			}
		}
	}

	outFileHandle.Close()

	return  true
}



//ff fe 22 = uft16 little endian bom
// basic control set  = 00 00  -> 00 7F this can be replaced to 00 -> 7F
// latin1 = 00 80 -> 00 FF this can be replaced to 80 -> FF


func fileExits( fp string ) bool {
	if _, err := os.Stat(fp); os.IsNotExist(err) {
		// path/to/whatever does not exist
		return  false
	}else {
		return true
	}
}


func IsUtf16( fp string ) bool {
	d,e := ioutil.ReadFile(Infile)

	if e != nil {
		fmt.Println(e)
		panic (e)
	}

	//fmt.Println(d);
	if d[0] == 255 && d[1] == 254 {
		return true
	}
	return false
}
