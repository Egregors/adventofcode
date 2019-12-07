module Main
  ( main
  )
where

import qualified Data.Set                      as Set

main :: IO ()
main = undefined
    -- start <- getCurrentTime
    -- print $ snd $ f 5 code
    -- end <- getCurrentTime
    -- print (diffUTCTime end start)

runCmd :: [Int] -> [Int] -> Int
runCmd input cmd = head . snd $ opApply 0 input (cmd, [])

-- opApply currentPosition input (xs, output) = ...
opApply :: Int -> [Int] -> ([Int], [Int]) -> ([Int], [Int])
opApply position input xss = f op input xss
 where
  opcode = normalizeOpcode $ show (fst xss !! position)
  [arg3Mode, arg2Mode, arg1Mode, _, op] = opcode
  arg1   = position + 1
  arg2   = position + 2
  arg3   = position + 3

  -- a + b -> c
  f '1' input (xs, output) = opApply newCursorPosition input (newXs, output)
   where
    newCursorPosition =
      if newValPosition == position then position else position + 4
    newXs = take newValPosition xs ++ newVal : drop (newValPosition + 1) xs
    newVal         = a + b
    a              = if arg1Mode == '0' then xs !! (xs !! arg1) else xs !! arg1
    b              = if arg2Mode == '0' then xs !! (xs !! arg2) else xs !! arg2
    newValPosition = xs !! arg3

  -- a * b -> c
  f '2' input (xs, output) = opApply newCursorPosition input (newXs, output)
   where
    newCursorPosition =
      if newValPosition == position then position else position + 4
    newXs = take newValPosition xs ++ newVal : drop (newValPosition + 1) xs
    newVal         = a * b
    a              = if arg1Mode == '0' then xs !! (xs !! arg1) else xs !! arg1
    b              = if arg2Mode == '0' then xs !! (xs !! arg2) else xs !! arg2
    newValPosition = xs !! arg3

  -- imp -> a
  f '3' input (xs, output) = opApply newCursorPosition
                                     (tail input)
                                     (newXs, output)
   where
    newCursorPosition =
      if newValPosition == position then position else position + 2
    newXs = take newValPosition xs ++ head input : drop (newValPosition + 1) xs
    newValPosition = xs !! arg1

  -- a -> out
  f '4' input (xs, output) = opApply (position + 2) input (xs, newOutput)
    where newOutput = output ++ [xs !! (xs !! arg1)]

  -- a /= 0 -> mv b : _
  f '5' input (xs, output) = if a /= 0
    then opApply b input (xs, output)
    else opApply (position + 3) input (xs, output)
   where
    a = if arg1Mode == '0' then xs !! (xs !! arg1) else xs !! arg1
    b = if arg2Mode == '0' then xs !! (xs !! arg2) else xs !! arg2

  -- a == 0 -> mv b
  f '6' input (xs, output) = if a == 0
    then opApply b input (xs, output)
    else opApply (position + 3) input (xs, output)
   where
    a = if arg1Mode == '0' then xs !! (xs !! arg1) else xs !! arg1
    b = if arg2Mode == '0' then xs !! (xs !! arg2) else xs !! arg2

  -- a < b -> c(1) : c(0)
  f '7' input (xs, output) = opApply newCursorPosition input (newXs, output)
   where
    newXs = take newValPosition xs ++ newVal : drop ((xs !! arg3) + 1) xs
    newCursorPosition =
      if newValPosition == position then position else position + 4
    newValPosition = xs !! arg3
    newVal         = if a < b then 1 else 0
    a              = if arg1Mode == '0' then xs !! (xs !! arg1) else xs !! arg1
    b              = if arg2Mode == '0' then xs !! (xs !! arg2) else xs !! arg2

    -- a == b -> c(1) : c(0)
  f '8' input (xs, output) = opApply newCursorPosition input (newXs, output)
   where
    newXs = take (xs !! arg3) xs ++ newVal : drop ((xs !! arg3) + 1) xs
    newCursorPosition =
      if newValPosition == position then position else position + 4
    newValPosition = xs !! arg3
    newVal         = if a == b then 1 else 0
    a              = if arg1Mode == '0' then xs !! (xs !! arg1) else xs !! arg1
    b              = if arg2Mode == '0' then xs !! (xs !! arg2) else xs !! arg2

  f _ input xss = xss

  normalizeOpcode [a, b, c, d, e] = [a, b, c, d, e]
  normalizeOpcode n               = normalizeOpcode $ '0' : n


-- p1

testCode :: [Int]
testCode =
  [ 3
  , 31
  , 3
  , 32
  , 1002
  , 32
  , 10
  , 32
  , 1001
  , 31
  , -2
  , 31
  , 1007
  , 31
  , 0
  , 33
  , 1002
  , 33
  , 7
  , 33
  , 1
  , 33
  , 31
  , 31
  , 1
  , 32
  , 31
  , 31
  , 4
  , 31
  , 99
  , 0
  , 0
  , 0
  ]

code :: [Int]
code =
  [ 3
  , 8
  , 1001
  , 8
  , 10
  , 8
  , 105
  , 1
  , 0
  , 0
  , 21
  , 34
  , 59
  , 76
  , 101
  , 114
  , 195
  , 276
  , 357
  , 438
  , 99999
  , 3
  , 9
  , 1001
  , 9
  , 4
  , 9
  , 1002
  , 9
  , 4
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 102
  , 4
  , 9
  , 9
  , 101
  , 2
  , 9
  , 9
  , 102
  , 4
  , 9
  , 9
  , 1001
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 101
  , 4
  , 9
  , 9
  , 102
  , 5
  , 9
  , 9
  , 101
  , 5
  , 9
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 1001
  , 9
  , 4
  , 9
  , 102
  , 4
  , 9
  , 9
  , 1001
  , 9
  , 4
  , 9
  , 1002
  , 9
  , 3
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 1002
  , 9
  , 3
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 1
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 1
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 1
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 1
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 1
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 1
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 1
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 1
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 1
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 102
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 99
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 1
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 1
  , 9
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1002
  , 9
  , 2
  , 9
  , 4
  , 9
  , 3
  , 9
  , 1001
  , 9
  , 1
  , 9
  , 4
  , 9
  , 3
  , 9
  , 101
  , 2
  , 9
  , 9
  , 4
  , 9
  , 99
  ]

allPhases :: [(Int, Int, Int, Int, Int)]
allPhases = filter hasDbls phases
 where
  phases =
    [ (a, b, c, d, e)
    | a <- [0 .. 4]
    , b <- [0 .. 4]
    , c <- [0 .. 4]
    , d <- [0 .. 4]
    , e <- [0 .. 4]
    ]

  hasDbls :: (Int, Int, Int, Int, Int) -> Bool
  hasDbls (a, b, c, d, e) = length set == length list
   where
    set  = Set.fromList list
    list = [a, b, c, d, e]

findThrusters instruction (ph0, ph1, ph2, ph3, ph4) = e
 where
  a = runCmd [ph0, 0] instruction
  b = runCmd [ph1, a] instruction
  c = runCmd [ph2, b] instruction
  d = runCmd [ph3, c] instruction
  e = runCmd [ph4, d] instruction

findMaxThruster = maximum . map (findThrusters code) $ allPhases


