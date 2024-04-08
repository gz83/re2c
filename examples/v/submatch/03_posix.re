// re2v $INPUT -o $OUTPUT

// Maximum number of capturing groups among all rules.
/*!maxnmatch:re2c*/

struct SemVer {
    major int
    minor int
    patch int
}

fn s2n(s string) int { // convert pre-parsed string to number
    mut n := 0
    for c in s { n = n * 10 + int(c - 48) }
    return n
}

fn parse(str string) ?SemVer {
    mut cur, mut mar := 0, 0

    // Allocate memory for capturing parentheses (twice the number of groups).
    mut yypmatch := []int{len: yymaxnmatch * 2}
    mut yynmatch := 0

    // Autogenerated tag variables used by the lexer to track tag values.
    /*!stags:re2c format = '\tmut @@ := 0\n'; */

    /*!re2c
        re2c:yyfill:enable = 0;
        re2c:define:YYCTYPE     = u8;
        re2c:define:YYPEEK      = "str[cur]";
        re2c:define:YYSKIP      = "cur += 1";
        re2c:define:YYBACKUP    = "mar = cur";
        re2c:define:YYRESTORE   = "cur = mar";
        re2c:define:YYSTAGP     = "@@{tag} = cur";
        re2c:define:YYSTAGN     = "@@{tag} = -1";
        re2c:define:YYSHIFTSTAG = "@@{tag} += @@{shift}";
        re2c:posix-captures = 1;

        num = [0-9]+;

        (num) "." (num) ("." num)? [\x00] {
            // `yynmatch` is the number of capturing groups
            if yynmatch != 4 { panic("expected 4 submatch groups") }

            // Even `yypmatch` values are for opening parentheses, odd values
            // are for closing parentheses, the first group is the whole match.
            return SemVer {
                major: s2n(str[yypmatch[2]..yypmatch[3]]),
                minor: s2n(str[yypmatch[4]..yypmatch[5]]),
                patch: if yypmatch[6] == -1 { 0 } else { s2n(str[yypmatch[6] + 1..yypmatch[7]]) }
            }
        }
        * { return none }
    */
}

fn main() {
    test := fn (result ?SemVer, expect ?SemVer) {
        if r := result {
            if e := expect { if r != e { panic("expected $e, got $r") } }
        } else {
            if _ := result { panic("expected none") }
        }
    }
    test(parse("23.34\0"), SemVer{23, 34, 0})
    test(parse("1.2.9999\0"), SemVer{1, 2, 9999})
    test(parse("1.a\0"), none)
}
