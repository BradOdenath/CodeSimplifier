package main

import (
	"strings"
	"unicode"
	"os"
	"bufio"
)

func main() {

}

//Representation of a java statement. Line number for record.
type javaStatement struct {
	statement string
	line int64
}

//Conventional Way to create a javaStatement
func createJavaStatement(str string, line int64) javaStatement {
	str = removeWhiteSpace(str)
	return javaStatement{str, line }
}

//Organize java statements into a slice.
type javaStatementList struct {
	statements[] javaStatement
}

func grabFile(fileName string) *File {
	file, err := os.Open(fileName)
	check(err)
	defer file.Close()
	return file
}

//Parse statements from a java file into a list.
func _parseFileToJavaStatements(fileName string) javaStatementList {
	/*
	TODO: This list for statement condensing.  That will be fun :D
	* -constants	NOT	INCLUDE
	* -statements		INCLUDE
	* -constructors		INCLUDE
	* -methods			INCLUDE
	*/
	file := grabFile(fileName)
	jsl := javaStatementList{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		scanStr := removeWhiteSpace(scanner.Text())
		tempRune := []rune(scanStr)
		if tempRune[(len(scanstr)-1)] == ';' {
			jsl.statements = append(jsl.statements,
				javaStatement{scanStr, -1})
		} else {
			//TODO: Recursion
		}
	}
}


//Remove white space from a java statement.
// Credit: Stackoverflow: Tim Cooper
func removeWhiteSpace(str string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, str)
}