module Main where

import Html

import Text
import Graphics.Element as Element
import Graphics.Collage as Collage
import Mouse
import Signal.Extra as SE
import Color
import Window
import Transform2D


--main : Signal Html.Html
--main =
--    Signal.constant (Html.div [] [Html.text "Hi", Html.fromElement (viewApp graph)])

main = Signal.map2 (pan (viewApp graph)) panPos Window.dimensions

mouseDelta : Signal (Int, Int)
mouseDelta =
    Signal.map (\((x2, y2),(x, y)) -> (x-x2, y-y2)) (SE.deltas Mouse.position)


dragDelta : Signal (Int, Int)
dragDelta =
    SE.keepWhen Mouse.isDown (0, 0) mouseDelta


dragPosition : Signal (Int, Int)
dragPosition =
    Signal.foldp (\ (x, y) (x2, y2) -> (x+x2, y+y2)) (0, 0) dragDelta

panPos : Signal (Int, Int)
panPos =
    Signal.map (\(x, y) -> (x, -y)) dragPosition

pan : Collage.Form -> (Int, Int) -> (Int, Int) -> Element.Element
pan f (x, y) (w, h) =
    Collage.collage w h [box w h, Collage.move (toFloat x, toFloat y) f]

box : Int -> Int -> Collage.Form
box w h =
    Collage.outlined (Collage.solid Color.black) (Collage.rect (Basics.toFloat w) (Basics.toFloat h))

viewApp : Graph -> Collage.Form
viewApp g =
    Collage.group (List.map viewNode g.nodes)

viewNode : Node -> Collage.Form
viewNode n =
    let
        (x, y) = n.position
        inputs = Collage.toForm (Element.flow Element.down (List.map viewPort n.process.inputs))
        outputs = Collage.toForm (Element.flow Element.down (List.map viewPort n.process.outputs))
    in
        Collage.groupTransform (Transform2D.translation x y)
            [ formFromString n.process.name
            , Collage.outlined (Collage.solid Color.black) (Collage.rect 150 50)
            , Collage.moveX -75 inputs
            , Collage.moveX 75 outputs
            ]

viewPort : Port -> Element.Element
viewPort p =
    Element.flow Element.down
        [ Element.centered (Text.fromString p.typ)
        ]


formFromString : String -> Collage.Form
formFromString =
    Text.fromString >> Element.centered >> Collage.toForm

type alias Node =
    { process: Process
    , position: (Float, Float)
    }

type alias Graph =
    { nodes: List Node
    }


type alias Port =
    { typ: String
    }


type alias Process =
    { name: String
    , inputs: List Port
    , outputs: List Port
    }

graph : Graph
graph =
    { nodes=
        [ { process= stringCLI
          , position=(-100, 0)
          }
        , { process= upper
          , position=(100, 0)
          }
        ]
    }

upper : Process
upper =
    { name="strings.ToUpper"
    , inputs= [{typ="string"}]
    , outputs= [{typ="string"}]
    }

stringCLI : Process
stringCLI =
    { name = "base.StringCLI"
    , inputs= []
    , outputs= [{typ="string"}]
    }