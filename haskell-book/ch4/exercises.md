
## Mood Swing

1. `Mood`
2. `Blah` or `Woot`
3. It should be `changeMood :: Mood -> Mood`
4.  ```
    changeMood Blah = Woot
    changeMood _ = Blah
    ```
5. Got it right eventually!

## Find the Mistakes

1. Nope, `true` is not in scope.
2. Nope, `x = 6` is for defining the value x.
3. Yea.
4. No, `Merry` and `Happy` are out of scope and don't implement `Ord`.
5. Nope, `[Num]` and `[Char]` can not be concatenated.

## Chapter Exercises

1. `[a] -> Int`
2. a) 5
   b) 3
   c) 2
   d) 5
3. The second one fails because you can't divide `Int`s
4. `div 6 (length [1, 2, 3])`
5. `Bool` and `True`
6. `Bool` and `False`
7. a) Yes, `True`.
   b) No, because lists require elements to be of the same type.
   c) Yes, `5`.
   d) Yes, `False`.
   e) No, because `9` is not of type `Bool`.
8.  ```
    isPalindrome :: (Eq a) => [a] -> Bool
    isPalindrome x = (reverse x) == x
    ```
9.  ```
    myAbs :: Integer -> Integer
    myAbs a = if a > 0 then a else -a
    ```
10. ```
    f :: (a, b) -> (c, d) -> ((b, d), (a, c))
    f x y = (((snd x), (snd y)), ((fst x), (fst y)))
    ```

### Correcting Syntax

1.  ```
    x = (+)
    f xs = w `x` 1
            where w = length xs
    ```
2.  `\ x -> x`
3.  `\ (x:xs) -> x`
4.  `f (a, b) = a`

### Match the function name to their types

1. c
2. b
3. a
4. d
