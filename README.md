## pipen

Pipen enables you to stream input into a shell command/filter
and then get its stdout without having to directly write to a
temp file.

## Usage in code
```golang
import "github.com/odeke-em/pipen"

...
f, err := os.Open("/var/log/auth.log.*")
if err != nil {
    // Handle the error here
    return
}

stdout, err := pipen.StreamCommand(stdin, "cut", "-d\":\"", "-f1")
...
```

## Working examples

See directory "examples/*"
