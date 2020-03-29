package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/thinkofher/pazzz/engine"
)

const (
	maxPasswordLength = 32
	secretEnv         = "PAZZZSECRET"
)

var (
	secretFlag    string
	lengthFlag    int
	lowercaseFlag bool
	uppercaseFlag bool
	digitsFlag    bool
	symbolsFlag   bool
	versionFlag   bool
)

func init() {
	flag.StringVar(&secretFlag, "secret", "", "")
	flag.IntVar(&lengthFlag, "len", 8, "")
	flag.BoolVar(&lowercaseFlag, "l", false, "")
	flag.BoolVar(&uppercaseFlag, "u", false, "")
	flag.BoolVar(&digitsFlag, "d", false, "")
	flag.BoolVar(&symbolsFlag, "s", false, "")
	flag.BoolVar(&versionFlag, "v", false, "")
	flag.Usage = func() {
		fmt.Fprint(os.Stdout, usage)
	}
	flag.Parse()
}

func run() error {
	// Show version number.
	if versionFlag {
		fmt.Println(version)
		os.Exit(0)
		return nil
	}

	// Handle situation with no arguments.
	if flag.NArg() == 0 {
		fmt.Fprint(os.Stderr, usage)
		os.Exit(1)
		return nil
	} else if flag.NArg() < 2 { // Handle situation without site argument.
		return fmt.Errorf("You have to provide login and site as arguments")
	}

	// Handle login.
	login := flag.Args()[0]

	// Handle site address.
	site := flag.Args()[1]

	// Handle length.
	if lengthFlag > maxPasswordLength {
		return fmt.Errorf("Given length: %d is bigger than maximum length: %d",
			lengthFlag, maxPasswordLength)
	}
	if lengthFlag <= 0 {
		return fmt.Errorf("Wrong length: %d", lengthFlag)
	}

	// Handle secret.
	secret := secretFlag
	if secret == "" {
		// If user did not specify secret in flag,
		// retrieve secret from environment varible.
		secret = os.Getenv(secretEnv)
	}
	if secret == "" {
		// If user did not specify secret in both flag
		// and environment varible, return error.
		return fmt.Errorf("Given secret is empty")
	}

	// Generate salt.
	s := engine.Salt(login, site, lengthFlag)

	// Generate entropy from salt.
	e, err := engine.Entropy(s, []byte(secret))
	if err != nil {
		return err
	}

	// Handle rules flags.
	rules := engine.Rules(lowercaseFlag, uppercaseFlag, digitsFlag, symbolsFlag)

	// Generate password.
	password := engine.Pass(e, *rules, lengthFlag)

	// Show password to user and exit with success.
	fmt.Println(string(password))
	os.Exit(0)

	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "%s: %s. Checkout pazz -h for more help.\n",
			name, err.Error())
		os.Exit(1)
	}
}
