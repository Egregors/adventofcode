import qualified Data.Map                      as Map
import           Data.Maybe
import           Data.List                      ( unfoldr )
import           System.Environment

main :: IO ()
main = do
  args    <- getArgs
  content <- readFile (head args)
  let linesOfFiles = lines content
  let mp           = makeOrbs linesOfFiles
  let res          = go mp
  print res

type Orbs = Map.Map String String

makeOrbs :: [String] -> Orbs
makeOrbs xs = Map.fromList pairs
  where pairs = map (listToPair . separateByOrb) xs

listToPair [a, b] = (b, a)
separateByOrb = separateBy ')'

separateBy :: Eq a => a -> [a] -> [[a]]
separateBy chr = unfoldr sep
 where
  sep [] = Nothing
  sep l  = Just . fmap (drop 1) . break (== chr) $ l

go mp = sum . map (\x -> find $ Map.lookup x mp) $ Map.keys mp
 where
  find = findOrbsCount mp 0
  findOrbsCount :: Map.Map String String -> Int -> Maybe String -> Int
  findOrbsCount _ count Nothing               = count
  findOrbsCount m count (Just nextCenterName) = findOrbsCount m
                                                              (count + 1)
                                                              nextCenter
    where nextCenter = Map.lookup nextCenterName m





