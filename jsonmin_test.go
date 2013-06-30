/*
 * Copyright 2013, Teo Mrnjavac <teo@kde.org>
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this
 * software and associated documentation files (the "Software"), to deal in the Software
 * without restriction, including without limitation the rights to use, copy, modify,
 * merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons whom the Software is furnished to do so, subject to the following
 * conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies
 * or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
 * INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
 * PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
 * HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF
 * CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE
 * OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package jsonmin

import (
	"testing"
)

func Test_Minify(t *testing.T) {
	tests := map[string]string{
		"": "",
		`{"foo":"bar","bar":["baz","bum","zam"],"something":10,"else":20}`: `{"foo":"bar","bar":["baz","bum","zam"],"something":10,"else":20}`,
		`
{
    "foo": "bar",    // this is cool
    "bar": [
        "baz", "bum", "zam"
    ],
/* the rest of this document is just fluff
   in case you are interested. */
    "something": 10,
    "else": 20
}

/* NOTE: You can easily strip the whitespace and comments 
   from such a file with the JSON.minify() project hosted 
   here on github at http://github.com/getify/JSON.minify 
*/
		`: `{"foo":"bar","bar":["baz","bum","zam"],"something":10,"else":20}`,
		`
{"/*":"*/","//":"",/*"//"*/"/*/"://
"//"}

`: `{"/*":"*/","//":"","/*/":"//"}`,
		`/*
this is a 
multi line comment */{

"foo"
:
    "bar/*"// something
    ,    "b\"az":/*
something else */"blah"

}
`: `{"foo":"bar/*","b\"az":"blah"}`,
		`{"foo": "ba\"r//", "bar\\": "b\\\"a/*z", 
    "baz\\\\": /* yay */ "fo\\\\\"*/o" 
}
`: `{"foo":"ba\"r//","bar\\":"b\\\"a/*z","baz\\\\":"fo\\\\\"*/o"}`,
	}

	for input, output := range tests {
		if res, err := Minify(input, true); res != output || err != nil {
			t.Error("Minify did not work as expected.\nInput:\n" + input + "$\nOutput:\n" + res + "$\nExpected:\n" + output + "$")
		}
	}
}
