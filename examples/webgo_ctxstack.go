package main

import (
    "fmt"
    "github.com/hoisie/web.go"
    "github.com/ziutek/kasia.go"
)

type Ctx struct {
    A, B  string
}

type LocalCtx struct {
    B string
    C int
}

var (
    tpl *kasia.Template
    data = &Ctx{"Hello!", "not specified"}
)

func hello(web_ctx *web.Context, val string) {
    // Local data
    var ld *LocalCtx

    if len(val) > 0 {
        ld = &LocalCtx{val, len(val)}
    }

    // Rendering data
    err := tpl.Run(web_ctx, data, ld)
    if err != nil {
        fmt.Fprint(web_ctx, "%", err, "%")
    }
}

const tpl_txt = `
<html><body>
    $A<br>
    Parameter: $B
    $if C:
        , length: $C
    $end
</body></html>
`

func main() {
    // Main template
    tpl = kasia.New()
    err := tpl.Parse([]byte(tpl_txt))
    if err != nil {
        fmt.Println("Main template", err)
        return
    }

    // This example can work in strict mode
    tpl.Strict = true

    // Web.go
    web.Get("/(.*)", hello)
    web.Run("0.0.0.0:9999")
}

