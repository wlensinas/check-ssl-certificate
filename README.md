# Intro

This is a simple script in Go for check the ssl certificate expiration from a site.

# Dependencies

For build and execute the code, you need to install Go https://go.dev/doc/install

# Install

- Clone the project: `git clone git@github.com:wlensinas/check-ssl-certificate.git` or [download zip](https://github.com/wlensinas/check-ssl-certificate/archive/refs/heads/master.zip)
- Move to the folder `cd check-ssl-certificate`

## What is the difference between go build and make build?

By default, the produced binary file contains debugging information and the symbol table. This can bloat the size of the file. To reduce the file size, you can include additional flags during the build process to strip this information from the binary. The following command will reduce the binary size by approximately 30 percent: `make build` execute this command `go build -ldflags "-w -s"`

# How to use

On macOS or linux operating systems:

1. Update `sites.json`
2. Execute `make run`

```
URL	 Certificate Date	 Days until expiration	
www.google.com.ar	 2023-06-12T08:27:56Z	 62	
www.infobae.com	 2023-06-11T17:03:10Z	 62	
```

# Source

* FreeCodeCamp article in https://www.freecodecamp.org/news/how-to-validate-ssl-certificates-in-go/
* Different forms of build Go binary files: Black Hat Go https://www.amazon.com/-/es/Tom-Steele/dp/1593278659