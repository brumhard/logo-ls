package main

var osDependentTests = []struct {
	args     []string
	testFile string
	td       string
}{
	{args: []string{"-1sh"}, testFile: "logo-ls-sh_darwin.snap", td: "Testing -sh (human readable size) execution"},
	{args: []string{"-1shRa"}, testFile: "logo-ls-shRa_darwin.snap", td: "Testing -shRa execution"},
}
