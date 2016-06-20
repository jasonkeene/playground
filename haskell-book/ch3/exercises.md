
## Exercises: Scope

1. Yes
2. No
3. No
4. Yes

## Exercises: Syntax Errors

1. No, unless you put parens around the infix operator then the output
    should be:
        [1, 2, 3, 4, 5, 6]
2. No, single quote is used for Char literals not for String. If you change it
   to use double quotes the output should be:
        "<3 Haskell"
3. Yes, the output should be:
        "<3 Haskell"

## Chapter Exercises

### Reading syntax

1. a) Valid
   b) Invalid
   c) Valid
   d) Invalid
   e) Invalid
   f) Valid
   g) Invalid
   g) Valid

2. a) [6, 12, 18] (d)
   b) "rainbow" (c)
   c) 10 (e)
   d) "Jules" (a)
   e) [2, 3, 5, 6, 8, 9] (b)

### Building functions

1. a) f x = x ++ "!"
   b) f x = [x !! 4]
   c) f x = drop 9 (x ++ "!")

2. a)   f :: String
        f x = x ++ "!"
   b)   g :: String
        g x = [x !! 4]
   c)   h :: String
        h x = drop 9 (x ++ "!")

3.  f :: String -> Char
    f x = x !! 2

4.  f :: Int -> Char
    f x = "Curry is awesome!" !! x

5.  rvrs :: String -> String
    rvrs x = drop 9 x ++ " " ++ take 2 (drop 6 x) ++ " " ++ take 5 x

6.  module Reverse where

    rvrs :: String -> String
    rvrs x = drop 9 x ++ " " ++ take 2 (drop 6 x) ++ " " ++ take 5 x
    
    main :: IO ()
    main = print $ rvrs "Curry is awesome"
