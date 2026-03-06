// This program is copyright 2022-2026 Percona LLC and/or its affiliates.
//
// THIS PROGRAM IS PROVIDED "AS IS" AND WITHOUT ANY EXPRESS OR IMPLIED
// WARRANTIES, INCLUDING, WITHOUT LIMITATION, THE IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE.
//
// This program is free software; you can redistribute it and/or modify it under
// the terms of the GNU General Public License as published by the Free Software
// Foundation, version 2.
//
// You should have received a copy of the GNU General Public License, version 2
// along with this program; if not, see <https://www.gnu.org/licenses/>.

package main

import (
	"os/exec"
	"regexp"
	"strings"
	"testing"
)

// TestConnectionFailure verifies that the tool reports an error when the MongoDB connection fails.
func TestConnectionFailure(t *testing.T) {
	// Use a port that is not listening to simulate a connection failure.
	badURI := "mongodb://127.0.0.1:19999"
	cmd := exec.Command("../../../bin/"+toolname, "--mongodb.uri", badURI, "check-all", "--all-databases")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Errorf("expected %s to exit with an error when MongoDB is unreachable, but it exited successfully", toolname)
	}
	want := "Cannot connect to the database"
	if !strings.Contains(string(out), want) {
		t.Errorf("expected output to contain %q when MongoDB is unreachable, got: %q", want, string(out))
	}
}

/*
Option --version
*/
func TestVersionOption(t *testing.T) {
	out, err := exec.Command("../../../bin/"+toolname, "--version").Output()
	if err != nil {
		t.Errorf("error executing %s --version: %s", toolname, err.Error())
	}
	// We are using MustCompile here, because hard-coded RE should not fail
	re := regexp.MustCompile(toolname + `\n.*Version v?\d+\.\d+\.\d+\n`)
	if !re.Match(out) {
		t.Errorf("%s --version returns wrong result:\n%s", toolname, out)
	}
}

func TestNoCommand(t *testing.T) {
	mockMongo := "mongodb://127.0.0.1:27017"
	out, err := exec.Command("../../../bin/"+toolname, "--mongodb.uri", mockMongo).Output()
	if err != nil {
		t.Errorf("error executing %s with no command: %s", toolname, err.Error())
	}

	want := "Usage: pt-mongodb-index-check show-help"
	if !strings.Contains(string(out), want) {
		t.Errorf("Output missmatch. Output %q should contain %q", string(out), want)
	}
}
