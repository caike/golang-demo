# ODevs Demo App

Here's two ways to run this program:

* `go run main.go "Message to be written to the file"`
* `go build` and `./odevs "Message to be written to the file"`


Then, check for the presence of a file called _file.txt_ which should
contain the message passed as argument to the program.

## Cross-Compiling for Windows64

In order to cross-compile this program for a Windows 64bit machine,
run the following:

`GOOS=windows GOARCH=amd64 go build -o odevs.exe`

This command generates a binary named **odevs.exe** which can be invoked
from Windows.
