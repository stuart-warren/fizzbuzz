package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/stuart-warren/fizzbuzz"
	"io"
	"log"
	"os"
)

// Run runs the command and returns nil if successful, otherwise
// an error.
func Run(in io.Reader, out io.Writer, args []string) error {
	flags := flag.NewFlagSet(args[0], flag.ExitOnError)
	flags.Parse(args[1:])
	_ = bufio.NewScanner(in)
	//for s.Scan() {
	//out.Write(s.Bytes())
	for i := 0; i < 100; i++ {
		io.WriteString(out, fmt.Sprintf("%s\n", fizzbuzz.FizzBuzz(i)))
	}
	//}
	//if err := s.Err(); err != nil {
	//	return err
	//}

	return nil
}

func main() {
	if err := Run(os.Stdin, os.Stdout, os.Args); err != nil {
		log.Fatalln(err)
	}
}
