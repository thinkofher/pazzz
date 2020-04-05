pazzz
=====

**Pazzz** is a stateless unix password manager. Instead of remembering password for every  site, you just have to remember your login, site and your secret phrase. **Pazzz** will always generate the same safe password for the input data, that you will provide.

### Installation

#### Before

In order to install **pazzz**, you need to have installed `go 1.14` compiler on your machine, have properly set `$GOPATH` variable and `$GOPATH/bin` path added to your `$PATH` variable.

For example. If you are *user* and you are using linux machine with bash, you can add these lines below to your `.bashrc` file in your home directory.

```bash
export GOPATH=/home/user/go
export PATH=$GOPATH/bin:$PATH
```
You can read more about `$GOPATH` [here](https://github.com/golang/go/wiki/GOPATH).

#### Installing with go

If you meet the above requirements, you can simple paste below command into your shell:

    $ go get github.com/thinkofher/pazzz


### Example usage

Let assume that your want to create an account on *greatsite.com* with your email *user@email.com*. Your secret phrase is *s3cr3d*. You can generate your password as below.

    $ pazzz -secret s3cr3d user@email.com greatsite.com
    oCgAgIeJ

And here is your password. Every time you will provide above data to pazzz, he will generate the same password. Everything you have to remember is your login, site and secret, which you should avoid sharing with anyone.

You can also use some extra options like, length of password or additional symbols.

    $ pazzz -secret s3cr3d -len 20 -u -l -s  user@email.com greatsite.com
    rZ.vP*sD)iJ,wG~hV%kP

As you can see, password have changed. So now, you have to also remember flags you have used, if you want to recreate password. But it is safer to use.

You can always check how to use **pazzz** with help flag, as in the command below.

    $ pazzz -h

### Environmental variables

#### `$PAZZZSECRET`

You can set this variable to your secret phrase, instead of entering it as a flag. It will keep your secret away from your shell history. Remember to do not keep it in your public dot files, because it will make your passwords vulnerable. The good way of storing your `$PAZZZSECRET` variable is to export it from separate shell source file, that you don't share with anyone.

#### `$PAZZZLEN`

You can set this variable to override default length of **pazzz**.

#### `$PAZZZFLAGS`

You can set this variable to override default flags (**u**ppercase and **l**owercase). Seperate flags with comma. This example value: `"d,u,l,s"`, will set every flag to true.

### Development

PRs are welcome. Just fork this project, create separate branch with descriptive name, commit and open pull request.

### License

This project is licensed under [BSD 3-Clause](LICENSE).
