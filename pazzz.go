package main

const (
	name    = "pazzz"
	version = "0.1.0"
)

const usage = `Pazzz is a stateless unix password manager.

Usage:

        pazz [flags] [login] [site]

Arguments:

        login str
            login that you will use with password.
        site str
            site that you will use generated password with.

The flags are:

        -len n
            length of password. cannot be bigger than 32.
            Default value is 8.
        -secret str
            secret to generate password from. You can also use
            "PAZZZSECRET" environment variable to set secret.
            Do no share your secret with anyone.
        -d
            include digits in password.
        -l
            include lowercase letters in password.
        -u
            include uppercase letters in password.
        -s
            include symbols in password.
        -v
            prints pazzz version to stdout.
        -h
            prints this message to stdout.

Check out github.com/thinkofher/pazzz for more information.
Author: Beniamin Dudek <beniamin.dudek@yahoo.com>.
`
