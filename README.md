# epp - Environment preprocessor

A small templating engine that allows you to use environmental variables. For
example:
```yaml
{% if HOME %}
...
{% endif %}
```

Will check if `$HOME` is set.

## Installation
```
$ go get github.com/soudy/epp
```

Or download the latest release.

## Usage
```
$ epp inputfile -o output
$ epp - < input > output
$ epp inputfile # Write to stdout
```
