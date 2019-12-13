module Main
  ( main
  )
where

import qualified Data.Set                      as Set
import           Debug.Trace

debug = flip trace 

main :: IO ()
main = undefined
    -- start <- getCurrentTime
    -- print $ snd $ f 5 code
    -- end <- getCurrentTime
    -- print (diffUTCTime end start)

runCmd :: [Int] -> [Int] -> ([Int], Int)
runCmd input cmd = (instruction, head output)
  where (instruction, output) = opApply 0 input (cmd, [])

-- opApply currentPosition input (xs, output) = ... | input is stack now
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

  f _ input xss = undefined

  normalizeOpcode [a, b, c, d, e] = [a, b, c, d, e]
  normalizeOpcode n               = normalizeOpcode $ '0' : n


-- p1

allPhases :: Int -> Int -> [(Int, Int, Int, Int, Int)]
allPhases from to = filter hasDbls phases
 where
  phases =
    [ (a, b, c, d, e)
    | a <- [from .. to]
    , b <- [from .. to]
    , c <- [from .. to]
    , d <- [from .. to]
    , e <- [from .. to]
    ]

  hasDbls :: (Int, Int, Int, Int, Int) -> Bool
  hasDbls (a, b, c, d, e) = length set == length list
   where
    set  = Set.fromList list
    list = [a, b, c, d, e]

findThrusters :: [Int] -> (Int, Int, Int, Int, Int) -> Int
findThrusters instruction (ph0, ph1, ph2, ph3, ph4) = loop
  e
  ( newInstructionForA
  , newInstructionForB
  , newInstructionForC
  , newInstructionForD
  , newInstructionForE
  )
 where
  (newInstructionForA, a) = runCmd [ph0, 0] instruction
  (newInstructionForB, b) = 
    if a /= 0 then runCmd [ph1, a] instruction else undefined
  (newInstructionForC, c) =
    if b /= 0 then runCmd [ph2, b] instruction else undefined
  (newInstructionForD, d) =
    if c /= 0 then runCmd [ph3, c] instruction else undefined
  (newInstructionForE, e) =
    if d /= 0 then runCmd [ph4, d] instruction else undefined

  loop :: Int -> ([Int], [Int], [Int], [Int], [Int]) -> Int
  loop signal (s1, s2, s3, s4, s5) = loop newSignal
                                          (newS1, newS2, newS3, newS4, newS5)
   where
    (newS1, a        ) = runCmd [signal] s1
    (newS2, b        ) 
        | a /= 0 = runCmd [a] s2 `debug` show a
        | otherwise = undefined
    (newS3, c        ) = if b /= 0 then runCmd [b] s3 else undefined
    (newS4, d        ) = if c /= 0 then runCmd [c] s4 else undefined
    (newS5, newSignal) = if d /= 0 then runCmd [d] s5 else undefined



-- p1 a = 0 b = 4, p2 a = 5 b = 9
-- findMaxThruster a b = maximum . map (findThrusters code 0) $ allPhases a b

-- p2

-- 18216
testCode :: [Int]
testCode =
  [ 3
  , 52
  , 1001
  , 52
  , -5
  , 52
  , 3
  , 53
  , 1
  , 52
  , 56
  , 54
  , 1007
  , 54
  , 5
  , 55
  , 1005
  , 55
  , 26
  , 1001
  , 54
  , -5
  , 54
  , 1105
  , 1
  , 12
  , 1
  , 53
  , 54
  , 53
  , 1008
  , 54
  , 0
  , 55
  , 1001
  , 55
  , 1
  , 55
  , 2
  , 53
  , 55
  , 53
  , 4
  , 53
  , 1001
  , 56
  , -1
  , 56
  , 1005
  , 56
  , 6
  , 99
  , 0
  , 0
  , 0
  , 0
  , 10
  ]

