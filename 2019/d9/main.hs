module Main
  ( main
  )
where
import           Debug.Trace

main :: IO ()
main = return ()

type RelBase = Int
type Position = Int
type Input = [Int]
type Instruction = [Int]
type RunTime = (Instruction, [Int])

runCmd :: Input -> Instruction -> RunTime
runCmd input cmd = opApply 0 input 0 (cmd, [])

-- opApply relBase currentPosition input (xs, output) = ...
opApply :: Position -> Input -> RelBase -> RunTime -> RunTime
opApply position input relBase xss = f op input xss
 where
  opcode = normalizeOpcode $ show (fst xss !! position)
  [arg3Mode, arg2Mode, arg1Mode, op1, op2] = opcode
  op     = [op1, op2]
  arg1   = position + 1
  arg2   = position + 2
  arg3   = position + 3

  normalizeOpcode [a, b, c, d, e] = [a, b, c, d, e]
  normalizeOpcode n               = normalizeOpcode $ '0' : n

  getValByMode :: Int -> Char -> [Int] -> Int
  getValByMode arg '0' xs = xs !! (xs !! arg) -- position mode 
  getValByMode arg '1' xs = xs !! arg -- immediate mode
  getValByMode arg '2' xs = xs !! (xs !! arg + relBase) -- relative mode

  -- a + b -> c
  f "01" input (xs, output) = opApply newCursorPosition
                                      input
                                      relBase
                                      (newXs, output)
   where
    newCursorPosition =
      if newValPosition == position then position else position + 4
    newXs = take newValPosition xs ++ newVal : drop (newValPosition + 1) xs
    newVal         = a + b
    a              = getValByMode arg1 arg1Mode xs
    b              = getValByMode arg2 arg2Mode xs
    newValPosition = xs !! arg3

  -- a * b -> c
  f "02" input (xs, output) = opApply newCursorPosition
                                      input
                                      relBase
                                      (newXs, output)
   where
    newCursorPosition =
      if newValPosition == position then position else position + 4
    newXs = take newValPosition xs ++ newVal : drop (newValPosition + 2) xs
    newVal         = a * b
    a              = getValByMode arg1 arg1Mode xs
    b              = getValByMode arg2 arg2Mode xs
    newValPosition = xs !! arg3

  -- imp -> a
  f "03" input (xs, output) = opApply newCursorPosition
                                      (tail input)
                                      relBase
                                      (newXs, output)
   where
    newCursorPosition =
      if newValPosition == position then position else position + 2
    newXs = take newValPosition xs ++ head input : drop (newValPosition + 1) xs
    newValPosition = xs !! arg1

  -- a -> out
  f "04" input (xs, output) = opApply (position + 2)
                                      input
                                      relBase
                                      (xs, newOutput)
   where
    newOutput = output ++ [a]
    a         = getValByMode arg1 arg1Mode xs -- 99 wtf????

  -- a /= 0 -> mv b : _
  f "05" input (xs, output) = if a /= 0
    then opApply b input relBase (xs, output)
    else opApply (position + 3) input relBase (xs, output)
   where
    a = getValByMode arg1 arg1Mode xs
    b = getValByMode arg2 arg2Mode xs

  -- a == 0 -> mv b
  f "06" input (xs, output) = if a == 0
    then opApply b input relBase (xs, output)
    else opApply (position + 3) input relBase (xs, output)
   where
    a = getValByMode arg1 arg1Mode xs
    b = getValByMode arg2 arg2Mode xs

  -- a < b -> c(1) : c(0)
  f "07" input (xs, output) = opApply newCursorPosition
                                      input
                                      relBase
                                      (newXs, output)
   where
    newXs = take newValPosition xs ++ newVal : drop ((xs !! arg3) + 1) xs
    newCursorPosition =
      if newValPosition == position then position else position + 4
    newValPosition = xs !! arg3
    newVal         = if a < b then 1 else 0
    a              = getValByMode arg1 arg1Mode xs
    b              = getValByMode arg2 arg2Mode xs

  -- a == b -> c(1) : c(0)
  f "08" input (xs, output) = opApply newCursorPosition
                                      input
                                      relBase
                                      (newXs, output)
   where
    newXs = take (xs !! arg3) xs ++ newVal : drop ((xs !! arg3) + 1) xs
    newCursorPosition =
      if newValPosition == position then position else position + 4
    newValPosition = xs !! arg3
    newVal         = if a == b then 1 else 0
    a              = getValByMode arg1 arg1Mode xs
    b              = getValByMode arg2 arg2Mode xs

  -- a -> relBase
  f "09" input (xs, output) = opApply (position + 2)
                                      input
                                      newRelBase
                                      (xs, output)
   where
    a          = getValByMode arg1 arg1Mode xs
    newRelBase = relBase + a

  f "99" _ xss = xss
  f _    _ xss = xss
