module Main where

import Html

import Text
import Graphics.Element as Element
import Graphics.Collage as Collage
import Mouse
import Color
import Window
import Transform2D
import Stage


main = Signal.map2 frame (Stage.panned (Signal.constant (viewApp graph))) Window.dimensions

frame : Collage.Form -> (Int, Int) -> Element.Element
frame f (w, h) =
    Collage.collage w h [box w h, f]

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