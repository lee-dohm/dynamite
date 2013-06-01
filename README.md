# Dynamite

A more readable system of writing test assertions for [Go][golang].

## Syntax

```go
import (
    "dynamite"
    "testing"
)

e = dynamite.E{}

func TestFoo(t *testing.T) {
    foo := Foo{}

    e.Expect(t, foo).To(BeNil)
}
```

[golang]: http://www.golang.org
