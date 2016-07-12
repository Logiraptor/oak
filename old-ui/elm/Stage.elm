module Stage where

import Graphics.Element
import Graphics.Collage
import Signal.Extra as SE
import Mouse


panned : Signal Bool -> Signal Graphics.Collage.Form -> Signal Graphics.Collage.Form
panned pannable form =
    Signal.map2 Graphics.Collage.move (panPos pannable) form

panPos : Signal Bool -> Signal (Float, Float)
panPos pannable =
    let
        tof (x, y) = (toFloat x, toFloat y)
        sum (x, y) (x2, y2) = (x+x2, y+y2)
        diff (x, y) (x2, y2) = (x2-x, y2-y)
        invertY (x, y) = (x, -y)
        mouseDelta = Signal.map ((uncurry diff) >> tof >> invertY) (SE.deltas Mouse.position)
        draggable = Signal.map2 (&&) Mouse.isDown pannable
        dragDelta = SE.keepWhen draggable (0, 0) mouseDelta
    in
        Signal.foldp sum (0, 0) dragDelta

