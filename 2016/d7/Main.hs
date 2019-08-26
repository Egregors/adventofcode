-- https://adventofcode.com/2016/day/7
module Main where

import Data.List.Split
import Data.List

main :: IO()
main = interact $ show . getTLSCount . words

type WordPairs = (Int, String)

getTLSCount xs = length $ filter isLineTLS xs
getWords xs p = [ snd x | x <- getPairs $ splitOneOf "[]" xs, p $ fst x ]
getPairs x = zip [1..] x

-- "abcdefg" -> ["abcd", "bcde", "cdef", "defg"]
splitBy4 xs
  | length xs <= 4 = [xs]
  | otherwise = [h] ++ splitBy4 t
  where
    h = take 4 xs
    t = drop 1 xs

isWordTLS s
  | length s /= 4 = False
  | length g == 3 && head g == last g = True
  | otherwise = False
  where
    g = group s

-- ioxxoj[asdfgh]zxcvbn -> True
isLineTLS :: String -> Bool
isLineTLS s
  | any isWordTLS $ concat $ map splitBy4 $ getWords s even = False
  | any isWordTLS $ concat $ map splitBy4 $ getWords s odd = True
  | otherwise = False



