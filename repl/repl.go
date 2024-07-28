package repl

/*
the REPL (Read Eval Print Loop) reads input
then sends it to the interpreter for evaluation
prints out the result/output
starts again
*/

// Tokenizes Monkey source code and prints the tokens

import (
	"bufio"
	"fmt"
	"interpretergo/lexer"
	"interpretergo/token"
	"io"
)

const PROMPT = ">>"

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()

		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
