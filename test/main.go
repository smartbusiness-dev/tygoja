package main

import (
	"log"
	"os"

	"github.com/smartbusiness-dev/tygoja"
)

func main() {
	gen := tygoja.New(tygoja.Config{
		Packages: map[string][]string{
			"github.com/smartbusiness-dev/tygoja/test/a": {"*"},
			"github.com/smartbusiness-dev/tygoja/test/b": {"*"},
			"github.com/smartbusiness-dev/tygoja/test/c": {"Example2", "Handler"},
		},
		Heading:              `declare var $app: c.Handler;`,
		WithPackageFunctions: true,
		// enable if you want to be able to import them
		// StartModifier: "export",
	})

	result, err := gen.Generate()
	if err != nil {
		log.Fatal(err)
	}

	if err := os.WriteFile("./types.d.ts", []byte(result), 0644); err != nil {
		log.Fatal(err)
	}

	// run `npx typedoc` to generate HTML docs from the above declarations
}
