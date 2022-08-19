package main

import "testing"

func TestSysinfo(t *testing.T) {
	got := Sysinfo("localhost", "0.0.a")
	want := "<h1>Hi</h1>\n<p>My hostname is: localhost</p>\n<p>My version is: 0.0.a</p>\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
