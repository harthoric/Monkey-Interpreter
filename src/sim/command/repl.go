package repl

import (
	"Simplex-Simia/evaluator"
	"Simplex-Simia/lexer"
	"Simplex-Simia/object"
	"Simplex-Simia/parser"
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os/user"
	"strings"
)

// starts the repl (read, evaluate, print, loop)
func Start(in io.Reader, out io.Writer, run bool, inpFile string) {

	// gets the current user
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	// used to store the 'PROMPT', when the user runs the repl (read, evaluate, print, loop)
	PROMPT := strings.Split(user.Name, " ")[0] + ">>"
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()
	macroEnv := object.NewEnvironment()

	if run {
		file, err := ioutil.ReadFile(inpFile)
		if err != nil {
			panic(err)
		}

		l := lexer.New(string(file))
		// parse this result through the parser
		p := parser.New(l)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
		}

		evaluated := evaluator.Eval(program, env)
		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	} else {
		for {
			fmt.Printf(PROMPT)
			// scan each line
			scanned := scanner.Scan()
			if !scanned {
				return
			}

			line := scanner.Text()
			l := lexer.New(line)
			p := parser.New(l)

			program := p.ParseProgram()
			if len(p.Errors()) != 0 {
				printParserErrors(out, p.Errors())
				continue
			}

			evaluator.DefineMacros(program, macroEnv)
			expanded := evaluator.ExpandMacros(program, macroEnv)

			evaluated := evaluator.Eval(expanded, env)
			if evaluated != nil {
				io.WriteString(out, evaluated.Inspect())
				io.WriteString(out, "\n")
			}
		}
	}

}

// to appear when there are errors
const MONKEY_FACE = `            __,__
   .--.  .-"     "-.  .--.
  / .. \/  .-. .-.  \/ .. \
 | |  '|  /   Y   \  |'  | |
 | \   \  \ 0 | 0 /  /   / |
  \ '- ,\.-"""""""-./, -' /
   ''-' /_   ^ ^   _\ '-''
       |  \._   _./  |
       \   \ '~' /   /
        '._ '-=-' _.'
           '-----'
`

// returns the errors, when present
func printParserErrors(out io.Writer, errors []string) {
	io.WriteString(out, MONKEY_FACE)
	io.WriteString(out, "Woops! We ran into some monkey business here!\n")
	io.WriteString(out, " parser errors:\n")
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}
