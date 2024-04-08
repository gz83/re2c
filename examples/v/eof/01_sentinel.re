// re2v $INPUT -o $OUTPUT

// Expect a null-terminated string.
fn lex(str string) int {
    mut cur := 0
    mut count := 0

loop: /*!re2c
    re2c:yyfill:enable = 0;
    re2c:define:YYCTYPE = u8;
    re2c:define:YYPEEK  = "str[cur]";
    re2c:define:YYSKIP  = "cur += 1";

    *      { return -1 }
    [\x00] { return count }
    [a-z]+ { count += 1; unsafe { goto loop } }
    [ ]+   {  unsafe { goto loop } }
    */
}

fn main() {
    assert lex("\0") == 0
    assert lex("one two three\0") == 3
    assert lex("f0ur\0") == -1
}
