/* BSD 3-Clause License
*
*  Copyright (c) 2018, Kristofer D. Davidson
*  All rights reserved.
*
*  Redistribution and use in source and binary forms, with or without
*  modification, are permitted provided that the following conditions are met:
*
*  Redistributions of source code must retain the above copyright notice, this
*  list of conditions and the following disclaimer.
*
*  Redistributions in binary form must reproduce the above copyright notice,
*  this list of conditions and the following disclaimer in the documentation
*  and/or other materials provided with the distribution.
*
*  Neither the name of the copyright holder nor the names of its
*  contributors may be used to endorse or promote products derived from
*  this software without specific prior written permission.
*
*  THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
*  AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
*  IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE ARE
*  DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE LIABLE
*  FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR CONSEQUENTIAL
*  DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF SUBSTITUTE GOODS OR
*  SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS INTERRUPTION) HOWEVER
*  CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN CONTRACT, STRICT LIABILITY,
*  OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE) ARISING IN ANY WAY OUT OF THE USE
*  OF THIS SOFTWARE, EVEN IF ADVISED OF THE POSSIBILITY OF SUCH DAMAGE.
 */

// Package quikcalc is a simple cli calculator.
package main // import "github.com/kidoda/quikcalc"

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func init() {
	var envList []string
	for key, val := range envList {
		err := os.Setenv(key, val)
		if err != nil {
			log.Printf("Unable to set env vars: %s \n", err)
		}
	}
}

func main() {
	var interactive = flag.Bool("interactive", true, "Run quikcalc in interactive mode.")
	flag.Parse()

	// right now only interactive mode is supported.
	if interactive != true {
		log.Fatalln("This version of quikcalc only supports interactive mode!")
	}

	var Usage = string{
		`QuikCalc is a simple, easy to use cli calculator written in Go.

	Usage:
	    quikcalc [-i] [expression]

	Available Options:
	    -i, --interactive        Interactive mode
	    -h, --help 		     Help information (this screen)

	For extended information consult the man page: "man quikcalc"`}
}
