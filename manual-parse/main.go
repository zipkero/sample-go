package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}
	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUsage(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "이름을 입력하고 엔터키를 눌러주세요"
	fmt.Fprintf(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("이름을 입력하세요")
	}
	return name, nil
}

type config struct {
	numTimes   int
	printUsage bool
}

func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	if len(args) != 1 {
		return c, errors.New("매개변수의 개수가 잘못되었습니다")
	}
	if args[0] == "-h" || args[0] == "--help" {
		c.printUsage = true
		return c, nil
	}
	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		return c, err
	}
	c.numTimes = numTimes
	return c, nil
}

func validateArgs(c config) error {
	if !(c.numTimes > 0) {
		return errors.New("0보다 큰 수를 입력하세요")
	}
	return nil
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUsage {
		printUsage(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}
	greetUser(c, name, w)
	return nil
}

var usageString = fmt.Sprintf(`
Usage: %s <integer> [-h|--help]
`, os.Args[0])

func printUsage(w io.Writer) {
	fmt.Fprintf(w, usageString)
}

func greetUser(c config, name string, w io.Writer) {
	msg := fmt.Sprintf("안녕하세요. %s\n", name)
	for i := 0; i < c.numTimes; i++ {
		fmt.Fprintf(w, msg)
	}
}
