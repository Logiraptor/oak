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
import Application
import NodeRenderer

main = Signal.map2 frame (Stage.panned (Signal.constant (viewApp graph))) Window.dimensions

frame : Collage.Form -> (Int, Int) -> Element.Element
frame f (w, h) =
    Collage.collage w h [box w h, f]

box : Int -> Int -> Collage.Form
box w h =
    Collage.outlined (Collage.solid Color.black) (Collage.rect (Basics.toFloat w) (Basics.toFloat h))

viewApp : Application.Graph -> Collage.Form
viewApp g =
    Collage.group ((List.map NodeRenderer.viewNode g.nodes)++(List.map (NodeRenderer.viewEdge g) g.edges))

graph : Application.Graph
graph =
    { nodes=
        [ { process= stringCLI
          , position=(-100, 0)
          }
        , { process= upper
          , position=(100, 0)
          }
        , { process= complex
          , position=(100, 200)
          }
        ]
    , edges=
        [ { from=(0,0)
          , to=(1,0)
          },
          { from=(1,0)
          , to=(2,0)
          },
          { from=(1,0)
          , to=(2,1)
          },
          { from=(0,0)
          , to=(2,2)
          }
        ]
    }


complex : Application.Process
complex =
    { name="complex.Process"
    , inputs= [{typ="string"}, {typ="int"}, {typ="*http.Request"}]
    , outputs= [{typ="map[string]int"}, {typ="bool"}]
    }


upper : Application.Process
upper =
    { name="strings.ToUpper"
    , inputs= [{typ="string"}]
    , outputs= [{typ="string"}]
    }

stringCLI : Application.Process
stringCLI =
    { name = "base.StringCLI"
    , inputs= []
    , outputs= [{typ="string"}]
    }