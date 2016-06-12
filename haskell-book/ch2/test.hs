sayHello :: String -> IO ()
sayHello x = putStrLn ("Hello, " ++ x ++ "!")

triple :: Integer -> Integer
triple x = x * 3

z a b = let y = a
            x = b
        in
        2 * (x + y)
