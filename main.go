package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/thinkofher/pazzz/engine"
)

const (
	maxPasswordLength = 32
	defaultPassLen    = 8
	secretEnv         = "PAZZZSECRET"
	lenEnv            = "PAZZZLEN"
	flagEnv           = "PAZZZFLAGS"
	flagSep           = ","
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
	var passLen int
	var err error

	if env := os.Getenv(lenEnv); env != "" {
		passLen, err = parsePassLen(env)
		if err != nil {
			handleError(err)
		}
	} else {
		passLen = defaultPassLen
	}

	flag.StringVar(&secretFlag, "secret", "", "")
	flag.IntVar(&lengthFlag, "len", passLen, "")
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
	var rules *[]engine.Rule

	// Get flags from environment variable.
	rules, err = parseFlags(os.Getenv(flagEnv))
	if err != nil {
		return err
	}

	if lowercaseFlag || uppercaseFlag || digitsFlag || symbolsFlag {
		// If user decides to use flags directly, respect them first.
		rules = engine.Rules(lowercaseFlag, uppercaseFlag, digitsFlag, symbolsFlag)
	}

	// Generate password.
	password := engine.Pass(e, *rules, lengthFlag)

	// Show password to user and exit with success.
	fmt.Println(string(password))
	os.Exit(0)

	return nil
}

func main() {
	if err := run(); err != nil {
		handleError(err)
	}
}

func parseFlags(flags string) (*[]engine.Rule, error) {
	rules := map[string]bool{
		"l": false,
		"u": false,
		"d": false,
		"s": false,
		"":  false, // handle empty flags string
	}

	for _, v := range strings.Split(flags, flagSep) {
		_, ok := rules[v]
		if !ok {
			return nil, fmt.Errorf("your %s env is corrupted. there is no %s flag", flagEnv, v)
		}
		rules[v] = true
	}

	return engine.Rules(rules["l"], rules["u"], rules["d"], rules["s"]), nil
}

func parsePassLen(env string) (int, error) {
	passLen, err := strconv.Atoi(env)
	if err != nil {
		return 0, fmt.Errorf("your %s env is corrupted. cannot parse '%s'", lenEnv, env)
	}
	if passLen > maxPasswordLength {
		return 0, fmt.Errorf("your %s env is corrupted. password cannot be bigger than %d", lenEnv, maxPasswordLength)
	}
	return passLen, nil
}

func handleError(e error) {
	fmt.Fprintf(os.Stderr, "%s: %s. Checkout %s -h for more help.\n",
		name, e.Error(), name)
	os.Exit(1)
}
