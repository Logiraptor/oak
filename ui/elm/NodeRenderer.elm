module NodeRenderer where

import Application exposing(..)
import Html
import Graphics.Collage as Collage
import Graphics.Element as Element
import Transform2D
import Color
import Text

get n xs = List.head (List.drop n xs)

viewEdge : Graph -> Edge -> Collage.Form
viewEdge g e =
    let
        (i1, i2) = e.from
        (j1, j2) = e.to
        from = Maybe.withDefault {position=(0,0), process={name="", inputs=[], outputs=[]}} (get i1 g.nodes)
        to = Maybe.withDefault {position=(0,0), process={name="", inputs=[], outputs=[]}} (get j1 g.nodes)
        (fromX, fromY) = from.position
        (toX, toY) = to.position
    in
        path (fromX+75, fromY-(20*(toFloat i2))) (toX-75, toY-(20*(toFloat j2)))


path : (Float, Float) -> (Float, Float) -> Collage.Form
path (startX, startY) (endX, endY) =
    let
        midX = (endX + startX) / 2
        points =
            [ (startX, startY)
            , (midX, startY)
            , (midX, endY)
            , (endX, endY)
            ]
    in
        Collage.traced Collage.defaultLine (Collage.path points)


viewNode : Node -> Collage.Form
viewNode n =
    let
        (x, y) = n.position
        (inputs, iheight) = formFlow (List.map viewPort n.process.inputs)
        (outputs, oheight) = formFlow (List.map viewPort n.process.outputs)
        height = max oheight iheight
    in
        Collage.groupTransform (Transform2D.translation x y)
            [ formFromString n.process.name
            , Collage.moveY (-height/2) (Collage.outlined (Collage.solid Color.black) (Collage.rect 150 (height+50)))
            , Collage.moveX -75 inputs
            , Collage.moveX 75 outputs
            ]

viewPort : Port -> Collage.Form
viewPort p =
    Collage.group
        [ Collage.filled Color.blue (Collage.circle 5)
        ]

formFromString : String -> Collage.Form
formFromString =
    Text.fromString >> Element.centered >> Collage.toForm

formFlow : List Collage.Form -> (Collage.Form, number)
formFlow l =
    let
        reducer next (rest, height) = ((Collage.moveY (height-20) next)::rest, height-20)
        (output, height) = List.foldl reducer ([], 20) l
    in
        (Collage.group output, -height)

emptyForm : Collage.Form
emptyForm =
    Collage.toForm Element.empty