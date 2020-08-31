# Gorrect
Is your submission correct? But in go, so **Gorrect**!

This tool is used to easily verify, check and time [code4bytes](https://github.com/Coding-Web-Community/code4bytes) challenge submissions against a pre-made list of testcases. Because of a larger userbase participating in the *weekly* challenges it became a bigger hassle to individually test and examine each challenge. Therefore we built a tool to do that for us. In the future we would like to have users submit their submissions using an API of some sorts and have them automatically tested, graded and rewarded. This tool is a critical part of that flow.

There is a big security risk attached to blindly running foreign scripts, that's why I would like to work up to a point where docker containers are being launched in which we can individually run these foreign scripts completely isolated from the host system and internet. Though that is still a big WIP.

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
