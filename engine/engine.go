// Package engine contains backend for pazzz unix password
// manager with 1000 test cases that will ensure, every
// new pazzz version will be compatible with old passwords.
//
// The heart of pazz are sha256 and hmac algorithms.
package engine

import (
	"crypto/hmac"
	"crypto/sha256"
	"fmt"
)

// Rule represents rule to build password from.
type Rule string

const (
	// Lowercase rule contains all ascii lowercase letters.
	Lowercase Rule = "abcdefghijklmnopqrstuvwxyz"

	// Upperase rule contains all ascii uppercase letters.
	Upperase Rule = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

	// Digits rule contains all ascii digits.
	Digits Rule = "0123456789"

	// Symbols contains 34 selected ascii symbols.
	Symbols Rule = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
)

// At returns byte from Rule which is placed
// at index equal to given i modulo length of Rule.
func (r Rule) At(i byte) byte {
	return r[int(i)%len(r)]
}

// Salt returns unique bytes slice created from
// give login, site name and length of password to generate.
func Salt(login, site string, length int) []byte {
	res := fmt.Sprintf("%s:%s:%d", login, site, length)
	return []byte(res)
}

// Entropy generates slice of bytes from given salt and
// secret, which are also slices of bytes.
func Entropy(salt, secret []byte) ([]byte, error) {
	m := hmac.New(sha256.New, secret)
	_, err := m.Write(salt)
	if err != nil {
		return nil, err
	}

	return m.Sum(nil), nil
}

// Pass returns password with given length generated
// from given entropy and slice of rules.
func Pass(entropy []byte, rules []Rule, length int) []byte {
	var res []byte
	for i := 0; i < length; i++ {
		r := rules[i%len(rules)]
		res = append(res, r.At(entropy[i]))
	}
	return res
}

// Rules function generates slice of rules from given boolean flags.
// By default (value of every flag is false) uses only letters.
func Rules(lowercase, uppercase, digits, symbols bool) *[]Rule {
	// By default use only letters.
	ans := []Rule{Lowercase, Upperase}
	if lowercase || uppercase || digits || symbols {
		// If client specify at least one rule, clear slice with
		// rules and initialize new one.
		ans = []Rule{}
	}
	if lowercase {
		ans = append(ans, Lowercase)
	}
	if uppercase {
		ans = append(ans, Upperase)
	}
	if digits {
		ans = append(ans, Digits)
	}
	if symbols {
		ans = append(ans, Symbols)
	}
	return &ans
}
