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

/*
 * Based on JSON.minify() 0.1 by Kyle Simpson, MIT license
 * https://github.com/getify/JSON.minify
 */

package jsonmin

import (
	"regexp"
	"strings"
)

func Minify(json string, stripSpace bool) (string, error) {
	tokenizer := regexp.MustCompile(`"|(/\*)|(\*/)|(//)|\n|\r`)
	whitespaceRx := regexp.MustCompile(`[ \t\n\r]*`)

	var (
		inString            = false
		inMultilineComment  = false
		inSinglelineComment = false
		newStr              = []string{}
		from                = 0
	)

	for _, matchLoc := range tokenizer.FindAllStringIndex(json, -1) {
		if !inMultilineComment && !inSinglelineComment {
			tmp2 := json[from:matchLoc[0]]
			if !inString && stripSpace {
				tmp2 = whitespaceRx.ReplaceAllString(tmp2, "")
			}
			newStr = append(newStr, tmp2)
		}
		from = matchLoc[1] //end of current match

		match := json[matchLoc[0]:matchLoc[1]]

		if match == `"` && !inMultilineComment && !inSinglelineComment {
			slEscapeRx := regexp.MustCompile(`(\\)*$`)
			escaped := slEscapeRx.FindString(json[:matchLoc[0]])
			if !inString || escaped == "" || len(escaped)%2 == 0 {
				inString = !inString
			}
			from-- //catch the " we just found
		} else if match == `/*` && !inString && !inMultilineComment && !inSinglelineComment {
			inMultilineComment = true
		} else if match == `*/` && !inString && inMultilineComment && !inSinglelineComment {
			inMultilineComment = false
		} else if match == `//` && !inString && !inMultilineComment && !inSinglelineComment {
			inSinglelineComment = true
		} else if isEOL, err := regexp.MatchString(`\n|\r`, match); isEOL && err == nil && !inString && !inMultilineComment && inSinglelineComment {
			inSinglelineComment = false
		} else if isWhitespace, err := regexp.MatchString(`\n|\r| |\t`, match); !inMultilineComment && !inSinglelineComment &&
			((!isWhitespace && err == nil) || !stripSpace) {
			newStr = append(newStr, match)
		}
	}

	if stripSpace { //after we're done with the last match all whitespace is safe to remove
		newStr = append(newStr, whitespaceRx.ReplaceAllString(json[from:], ""))
	} else {
		newStr = append(newStr, json[from:])
	}

	return strings.Join(newStr, ""), nil
}
