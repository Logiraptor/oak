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
import Time
import SceneTree



graphSig = Signal.map changingGraph (Time.every Time.second)

changingGraph : Time.Time -> Application.Graph
changingGraph t =
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
          , to=(2,((round (Time.inSeconds t)) % 3))
          }
        ]
    }


main =
    (Signal.map2 onStuff Mouse.position graphSig)
        `withExtra`
        Signal.map2 frame (Stage.panned (Signal.constant True) (Signal.map viewApp graphSig)) Window.dimensions


onNode : (Int, Int) -> Application.Node -> (Bool, Bool)
onNode (x, y) n =
    let
        (fx, fy) = (toFloat x, toFloat y)
        (nx, ny) = n.position
    in
        (fx > nx && fx < (nx + 150),
        fy > ny && fy < (ny + (20 * (max (toFloat (List.length n.process.inputs)) (toFloat (List.length n.process.outputs))))))


onStuff : (Int, Int) -> Application.Graph -> List (Bool, Bool)
onStuff pos graph =
    List.map (onNode pos) graph.nodes


withExtra : Signal a -> Signal Element.Element -> Signal Element.Element
withExtra extra rest =
    Signal.map2 (Element.above << Element.show) extra rest



frame : Collage.Form -> (Int, Int) -> Element.Element
frame f (w, h) =
    Collage.collage w h [f]

box : Int -> Int -> Collage.Form
box w h =
    Collage.outlined (Collage.solid Color.black) (Collage.rect (Basics.toFloat w) (Basics.toFloat h))

viewApp : Application.Graph -> Collage.Form
viewApp g =
    Collage.group ((List.map NodeRenderer.viewNode g.nodes)++(List.map (NodeRenderer.viewEdge g) g.edges))



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