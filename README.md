# Gorrect
Is your submission correct? But in go, so **Gorrect**!

This tool is used to easily verify and check [code4bytes](https://github.com/Coding-Web-Community/code4bytes) challenges against a premade list of testcases.

Feel free to contribute, but before you do make an issue!

# What test cases should look like
```
{"args": ["vile", "evil"],"truth": "true"}
{"args": ["serbia", "rabies"],"truth": "true"}
{"args": ["mathis", "golang"], "truth": "false"}
```

# What should the file structure look like

```
main.go
test_cases.json
submissions/
    submission1.py
    submission2.py
```

# How do I run it

```
$ go run main.go
```
or
```
$ go build -o gorrect main.go
$ ./gorrect
```
