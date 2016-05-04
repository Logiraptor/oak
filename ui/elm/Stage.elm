module Stage where

import Graphics.Element
import Graphics.Collage
import Signal.Extra as SE
import Mouse


panned : Signal Graphics.Collage.Form -> Signal Graphics.Collage.Form
panned form =
    Signal.map2 Graphics.Collage.move panPos form

panPos : Signal (Float, Float)
panPos =
    let
        tof (x, y) = (toFloat x, toFloat y)
        sum (x, y) (x2, y2) = (x+x2, y+y2)
        diff (x, y) (x2, y2) = (x2-x, y2-y)
        invertY (x, y) = (x, -y)
        mouseDelta = Signal.map ((uncurry diff) >> tof >> invertY) (SE.deltas Mouse.position)
        dragDelta = SE.keepWhen Mouse.isDown (0, 0) mouseDelta
    in
        Signal.foldp sum (0, 0) dragDelta

