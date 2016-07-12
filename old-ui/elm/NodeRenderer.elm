module NodeRenderer (..) where

import Application exposing (..)
import Html
import Graphics.Collage as Collage
import Graphics.Element as Element
import Transform2D
import Color
import Text
import SceneTree


get n xs =
  List.head (List.drop n xs)


viewEdge : Graph -> Edge -> SceneTree.SceneTree
viewEdge g e =
  let
    ( i1, i2 ) =
      e.from

    ( j1, j2 ) =
      e.to

    from =
      Maybe.withDefault { position = ( 0, 0 ), process = { name = "", inputs = [], outputs = [] } } (get i1 g.nodes)

    to =
      Maybe.withDefault { position = ( 0, 0 ), process = { name = "", inputs = [], outputs = [] } } (get j1 g.nodes)

    ( fromX, fromY ) =
      from.position

    ( toX, toY ) =
      to.position
  in
    elbow_path ( fromX + 75, fromY - (20 * (toFloat i2)) ) ( toX - 75, toY - (20 * (toFloat j2)) )


elbow_path : ( Float, Float ) -> ( Float, Float ) -> SceneTree.SceneTree
elbow_path ( startX, startY ) ( endX, endY ) =
  let
    midX =
      (endX + startX) / 2

    points =
      [ ( startX, startY )
      , ( midX, startY )
      , ( midX, endY )
      , ( endX, endY )
      ]

    path =
      Collage.path points

    form =
      Collage.traced Collage.defaultLine path

    node =
      SceneTree.Leaf ( 0, 0 ) form
  in
    SceneTree.NoCollide node


viewNode : Node -> SceneTree.SceneTree
viewNode n =
  let
    ( x, y ) =
      n.position

    ( inputs, iheight ) =
      formFlow (List.map viewPort n.process.inputs)

    ( outputs, oheight ) =
      formFlow (List.map viewPort n.process.outputs)

    height =
      max oheight iheight
  in
    SceneTree.Move
      ( x, y )
      (SceneTree.Group
        [ SceneTree.Leaf ( 0, 0 ) (formFromString n.process.name)
        , SceneTree.Move ( 0, -height / 2 ) (SceneTree.Leaf ( 150, height + 50 ) (Collage.outlined (Collage.solid Color.black) (Collage.rect 150 (height + 50))))
        , SceneTree.Move ( -75, 0 ) (SceneTree.Leaf ( 0, 0 ) inputs)
        , SceneTree.Move ( 75, 0 ) (SceneTree.Leaf ( 0, 0 ) outputs)
        ]
      )


viewPort : Port -> Collage.Form
viewPort p =
  Collage.group
    [ Collage.filled Color.blue (Collage.circle 5)
    ]


formFromString : String -> Collage.Form
formFromString =
  Text.fromString >> Element.centered >> Collage.toForm


formFlow : List Collage.Form -> ( Collage.Form, number )
formFlow l =
  let
    reducer next ( rest, height ) =
      ( (Collage.moveY (height - 20) next) :: rest, height - 20 )

    ( output, height ) =
      List.foldl reducer ( [], 20 ) l
  in
    ( Collage.group output, -height )


emptyForm : Collage.Form
emptyForm =
  Collage.toForm Element.empty
