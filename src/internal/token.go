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

package quikcalc // import "github.com/kidoda/quikcalc"

import (
	"strconv"
)

type Token int

const (
	// Number Literals
	ZERO Token = iota

	ONE   // 1
	TWO   // 2
	THREE // 3
	FOUR  // 4
	FIVE  // 5
	SIX   // 6
	SEVEN // 7
	EIGHT // 8
	NINE  // 9

	// Operators
	ADD  // +
	SUB  // -
	MUL  // *
	DIV  // /
	MOD  // %
	EXP  // ^
	EVAL // =

	// MISC
	LPARANTHESIS // (
	RPARENTHESIS // )
	ENTER        // \CR

)

var tokens = [...]string{
	ZERO:  '0',
	ONE:   '1',
	TWO:   '2',
	THREE: '3',
	FOUR:  '4',
	FIVE:  '5',
	SIX:   '6',
	SEVEN: '7',
	EIGHT: '8',
	NINE:  '9',

	ADD:  '+',
	SUB:  '-',
	MUL:  '*',
	DIV:  '/',
	MOD:  '%',
	EXP:  '^',
	EVAL: '=',

	LPARANTHESIS: '(',
	RPARENTHESIS: ')',
}
