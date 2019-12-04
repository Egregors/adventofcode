import           Data.List

main :: IO ()
main = print
  (foldl (\count el -> if isValid el then count + 1 else count)
         0
         [100000 .. 999999]
  )

isValid num = isSorted' xs && hasPair xs
 where
  xs = show num
  hasPair ys = not (null pairs)
  pairs = filter (\x -> length x == 2) $ group xs
  isSorted' xs = and . zipWith (<=) xs $ tail xs
