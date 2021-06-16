package main

var osDependentTests = []struct {
	args     []string
	testFile string
	td       string
}{
	{args: []string{"-1sh"}, testFile: "logo-ls-sh_linux.snap", td: "Testing -sh (human readable size) execution"},
	{args: []string{"-1shRa"}, testFile: "logo-ls-shRa_linux.snap", td: "Testing -shRa execution"},
}