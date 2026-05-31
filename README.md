                                            ASCII-ART 
Ascii-art is a program that represent characters of a standard ascii in a graphical form and it's CLI (command line interface) based.
From the program we have a list of banners, banners are the text file that contains the graphical form of the ascii characters.

Below is a list of the banners:
  * shadow.txt
  * standard.txt
  * thinkertoy.txt
 
Package asciimaps
This package contains the following function.
  - Loadfile : This is a function that load a banner from a file into the program and return a slice a map.
  - RensderLine : The function parameters are string and a map whos value is a slice of string and a rune as the key, it's the function that builds the characters, depending on the range the characters in the banner and returns it in a slice of string form. 
  - SplitInput : As the name implies, it's use to split the input and return a slice of string.
  - ValidateInput : It's checks for the validity of the string, if the input is an ascii character or not, it returns a rune and an error.
  - GenerateArt : This is where the whole character is now formed and ready for display to the terminal, this function takes a string and a map whose values are slice of string and the keys are rune and returns strings.

Package main
  This package contains the main function which is the entery point of every go program and where all the execution is done. It's in this function that all the other function from the asciimap package are called for execution.
