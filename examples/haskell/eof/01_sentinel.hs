-- Generated by re2c
#1 "haskell/eof/01_sentinel.re"
-- re2hs $INPUT -o $OUTPUT
{-# LANGUAGE OverloadedStrings #-}
{-# LANGUAGE OverloadedRecordDot #-}

import Control.Monad (when)
import Data.ByteString as BS

data State = State {
    str :: BS.ByteString,
    cur :: Int
} deriving (Show)

-- expect a null-terminated string

#18 "haskell/eof/01_sentinel.hs"
yy0 :: State -> Int -> Int
yy0 _s _cnt =
    let yych = BS.index _s.str _s.cur in
    let _t = _s{cur = _s.cur + 1} in let _s = _t in
    case yych of
        _c | 0x00 == _c ->
            yy1 _s _cnt
        _c | 0x20 == _c ->
            yy3 _s _cnt
        _c | 0x61 <= _c && _c <= 0x7A ->
            yy5 _s _cnt
        _c | True ->
            yy2 _s _cnt

yy1 :: State -> Int -> Int
yy1 _s _cnt =
#22 "haskell/eof/01_sentinel.re"
    _cnt
#37 "haskell/eof/01_sentinel.hs"

yy2 :: State -> Int -> Int
yy2 _s _cnt =
#21 "haskell/eof/01_sentinel.re"
    (-1)
#43 "haskell/eof/01_sentinel.hs"

yy3 :: State -> Int -> Int
yy3 _s _cnt =
    let yych = BS.index _s.str _s.cur in
    case yych of
        _c | 0x20 == _c ->
            let _t = _s{cur = _s.cur + 1} in let _s = _t in
            yy3 _s _cnt
        _c | True ->
            yy4 _s _cnt

yy4 :: State -> Int -> Int
yy4 _s _cnt =
#24 "haskell/eof/01_sentinel.re"
    lexer _s _cnt
#59 "haskell/eof/01_sentinel.hs"

yy5 :: State -> Int -> Int
yy5 _s _cnt =
    let yych = BS.index _s.str _s.cur in
    case yych of
        _c | 0x61 <= _c && _c <= 0x7A ->
            let _t = _s{cur = _s.cur + 1} in let _s = _t in
            yy5 _s _cnt
        _c | True ->
            yy6 _s _cnt

yy6 :: State -> Int -> Int
yy6 _s _cnt =
#23 "haskell/eof/01_sentinel.re"
    lexer _s (_cnt + 1)
#75 "haskell/eof/01_sentinel.hs"

lexer :: State -> Int -> Int
lexer _s _cnt =
    yy0 _s _cnt

#25 "haskell/eof/01_sentinel.re"


main :: IO ()
main = do
    let test s n = when (lexer (State {str = s, cur = 0}) 0 /= n) $ error "failed"
    test "\0" 0
    test "one two three\0" 3
    test "f0ur\0" (-1)
