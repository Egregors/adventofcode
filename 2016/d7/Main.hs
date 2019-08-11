-- https://adventofcode.com/2016/day/7
module Main where

import Data.List.Split
import Data.List

main :: IO()
main = do
  let s = "abba[mnop]qrst"
  putStrLn $ intercalate " " $ getWords s

getWords :: String -> [String]
getWords s = splitOneOf "[]" s

isTLS :: String -> Bool
isTLS s = 
