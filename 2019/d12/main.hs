
module Main
  ( main
  )
where

import           Debug.Trace

main :: IO ()
main = return ()


data Point = Point {
    x :: Int
    , y :: Int
    , z :: Int
} deriving (Eq, Show)

data Moon = Moon {
    pos :: Point
    , vel :: Point
} deriving (Eq, Show)

-- <x=-1, y=0, z=2>
-- <x=2, y=-10, z=-7>
-- <x=4, y=-8, z=8>
-- <x=3, y=5, z=-1>

moon1 = Moon (Point (-1) 0 2) (Point 0 0 0)
moon2 = Moon (Point 2 (-10) (-7)) (Point 0 0 0)
moon3 = Moon (Point 4 (-8) 8) (Point 0 0 0)
moon4 = Moon (Point 3 5 (-1)) (Point 0 0 0)

moons = [moon1, moon2, moon3, moon4]

-- <x=-8, y=-10, z=0>
-- <x=5, y=5, z=10>
-- <x=2, y=-7, z=3>
-- <x=9, y=-8, z=-3>
moon1' = Moon (Point (-8) (-10) 0) (Point 0 0 0)
moon2' = Moon (Point 5 5 10) (Point 0 0 0)
moon3' = Moon (Point 2 (-7) 3) (Point 0 0 0)
moon4' = Moon (Point 9 (-8) (-3)) (Point 0 0 0)

moons' = [moon1', moon2', moon3', moon4']

-- <x=-10, y=-10, z=-13>
-- <x=5, y=5, z=-9>
-- <x=3, y=8, z=-16>
-- <x=1, y=3, z=-3>

moon1'' = Moon (Point (-10) (-10) (-13)) (Point 0 0 0)
moon2'' = Moon (Point 5 5 (-9)) (Point 0 0 0)
moon3'' = Moon (Point 3 8 (-16)) (Point 0 0 0)
moon4'' = Moon (Point 1 3 (-3)) (Point 0 0 0)

moons'' = [moon1'', moon2'', moon3'', moon4'']


pointToList :: Point -> [Int]
pointToList p = [x p, y p, z p]

getTotalEnergy :: Moon -> Int
getTotalEnergy m = sum (map abs [x p, y p, z p])
  * sum (map abs [x v, y v, z v])
 where
  p = pos m
  v = vel m

getTotalEnergyForAll :: [Moon] -> Int
getTotalEnergyForAll ms = sum $ map getTotalEnergy ms

getEnergyAfterN :: Int -> [Moon] -> Int
getEnergyAfterN i ms = getTotalEnergyForAll $ calc i ms

getVel :: (Moon, Moon) -> (Moon, Moon)
getVel (m1, m2) = (Moon (pos m1) newVel1, Moon (pos m2) newVel2)
 where
  [posM1, posM2] = [ [x p, y p, z p] | p <- map pos [m1, m2] ]
  [velM1, velM2] = [ [x p, y p, z p] | p <- map vel [m1, m2] ]
  points         = zip posM1 posM2
  comp a b | a < b     = 1
           | a > b     = -1
           | otherwise = 0
  velDiff1 = map (\(x1, x2) -> comp x1 x2) points
  velDiff2 = map (\(x1, x2) -> comp x2 x1) points
  newVel1  = Point a b c where [a, b, c] = zipWith (+) velM1 velDiff1
  newVel2  = Point a b c where [a, b, c] = zipWith (+) velM2 velDiff2

calcVel :: Moon -> [Moon] -> Moon
calcVel base [] = Moon newPos (vel base)
 where
  newPos    = Point a b c
  [a, b, c] = zipWith (+) (pointToList $ pos base) (pointToList $ vel base)
calcVel base moons = calcVel newBase ms
 where
  m : ms  = moons
  newBase = fst $ getVel (base, m)

calcVelForAll :: [Moon] -> [Moon]
calcVelForAll [m1, m2, m3, m4] = [newM1, newM2, newM3, newM4]
 where
  newM1 = calcVel m1 [m2, m3, m4]
  newM2 = calcVel m2 [m1, m3, m4]
  newM3 = calcVel m3 [m1, m2, m4]
  newM4 = calcVel m4 [m1, m2, m3]

calc :: Int -> [Moon] -> [Moon]
calc 0 ms = ms
calc i ms = calc (i - 1) (calcVelForAll ms)

