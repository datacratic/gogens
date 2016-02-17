// Copyright (c) 2016 Datacratic. All rights reserved.

package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var (
	pkg    = flag.String("package", "", "package name for which to include templates")
	folder = flag.String("folder", ".", "folder which contains the templates to be included")
	output = flag.String("output", "include_suffix.go", "output filename")
	suffix = flag.String("suffix", "", "suffix to include")
)

func main() {
	flag.Parse()
	fmt.Println("go generate include_suffix", *pkg, *folder, *output, *suffix)

	if *pkg == "" {
		flag.Usage()
		panic("package needs to be set")
		return
	}
	if *suffix == "" {
		flag.Usage()
		panic("suffix needs to be set")
		return
	}
	suffixUpper := strings.ToUpper(*suffix)

	fs, _ := ioutil.ReadDir(*folder)
	out, _ := os.Create(*output)

	out.WriteString("package " + *pkg + "\n\n")
	out.WriteString("//go:generate include_suffix")
	out.WriteString(" -package=" + *pkg)
	out.WriteString(" -folder=" + *folder)
	out.WriteString(" -output=" + *output)
	out.WriteString(" -suffix=" + *suffix)
	out.WriteString("\n\nconst(\n")

	vars := []string{}

	for _, f := range fs {
		if strings.HasSuffix(f.Name(), "."+*suffix) {
			varName := strings.TrimSuffix(f.Name(), "."+*suffix) + suffixUpper
			out.WriteString(varName + " = `")
			vars = append(vars, varName)
			f, err := os.Open(*folder + "/" + f.Name())
			if err != nil {
				fmt.Println(err)
			}
			io.Copy(out, f)
			out.WriteString("`\n")
		}
	}
	out.WriteString(")\n\n")

	out.WriteString("func Get" + suffixUpper + "() []string {\n\treturn []string{ ")
	for i, v := range vars {
		out.WriteString(v)
		if i < len(vars)-1 {
			out.WriteString(", ")
		}
	}
	out.WriteString(" }\n}")
}
