module Lib
    ( someFunc
    ) where

import System.Environment

someFunc :: IO ()
someFunc = do
  (x:xs) <- getArgs
  putFile x

putFile :: String -> IO ()
putFile f = do
  cs <- readFile f
  putStr cs
