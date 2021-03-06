// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

/*

Cov is a rudimentary code coverage tool.

Usage:
	go tool cov [-lsv] [-g substring] [-m minlines] [6.out args]

Given a command to run, it runs the command while tracking which
sections of code have been executed.  When the command finishes,
cov prints the line numbers of sections of code in the binary that
were not executed.   With no arguments it assumes the command "6.out".


The options are:

	-l
		print full path names instead of paths relative to the current directory
	-s
		show the source code that didn't execute, in addition to the line numbers.
	-v
		print debugging information during the run.
	-g substring
		restrict the coverage analysis to functions or files whose names contain substring
	-m minlines
		only report uncovered sections of code larger than minlines lines

The program is the same for all architectures: 386, amd64, and arm.

*/
package main
