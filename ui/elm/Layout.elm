module Layout exposing (..)

import Collage
import DataTypes exposing (..)
import Element
import Graph
import Html
import Html.Attributes
import Html.Events
import String
import Svg
import Svg.Attributes


-- Consts


procWidth =
    150


portSpacing =
    25


portY index =
    portSpacing + (toFloat index) * portSpacing


procHeight : Process -> Float
procHeight proc =
    portSpacing + (toFloat (Basics.max (List.length proc.type'.inputs) (List.length proc.type'.outputs)) * portSpacing)


layout : Model -> Html.Html Msg
layout model =
    Html.div [ Html.Events.onMouseUp Unclick ]
        [ navbar
        , Html.div [ Html.Attributes.class "content" ]
            [ sidebar, editor model ]
        ]


navbar : Html.Html Msg
navbar =
    Html.div [ Html.Attributes.class "navbar" ]
        [ Html.h1 [] [ Html.text "Oak Editor" ] ]


sidebar : Html.Html Msg
sidebar =
    Html.div [ Html.Attributes.class "sidebar" ]
        [ searchbar ]


searchbar : Html.Html Msg
searchbar =
    Html.input
        [ Html.Attributes.type' "text"
        , Html.Attributes.class "search"
        ]
        []


editor : Model -> Html.Html Msg
editor model =
    let
        elems =
            Graph.mapNodes viewNodeAndPipe model.graph
    in
        Html.div [ Html.Attributes.class "editor" ]
            [ Svg.svg
                [ Svg.Attributes.width ("100%")
                , Svg.Attributes.height ("100%")
                ]
                [ Svg.g [] elems
                , Svg.g [] defs
                  --, Svg.g []
                  --    [ text 14 [] (toString model)
                  --    ]
                ]
            ]


viewNodeAndPipe : Graph.NodeContext Process Pipe -> Svg.Svg Msg
viewNodeAndPipe ctx =
    let
        box =
            viewNode ctx
    in
        box


viewPipe : Graph.NodeContext Process Pipe -> ( Graph.Edge Pipe, Graph.Node Process ) -> Svg.Svg Msg
viewPipe ctx ( pipe, to ) =
    let
        ( x, y ) =
            ctx.node.label.pos

        ( x2, y2 ) =
            to.label.pos
    in
        line ( procWidth, portY pipe.label.output )
            ( x2 - x, (y2 - y) + (portY pipe.label.input) )
            [ Svg.Attributes.stroke "#000"
            , Svg.Attributes.fill "none"
            , Svg.Attributes.markerEnd "url(#arrow-head)"
            ]
            []


line : ( Float, Float ) -> ( Float, Float ) -> List (Svg.Attribute Msg) -> List (Svg.Svg Msg) -> Svg.Svg Msg
line from to =
    let
        pre p x y =
            p ++ (toString x) ++ " " ++ (toString y)

        ( x1, y1 ) =
            from

        ( x2, y2 ) =
            to

        d =
            String.join ", "
                [ pre "M " x1 y1
                , pre "C " (x1 + 100) y1
                , pre "" (x2 - 100) y2
                , pre "" x2 y2
                ]
    in
        Svg.path << (::) (Svg.Attributes.d d)


viewNode : Graph.NodeContext Process Pipe -> Svg.Svg Msg
viewNode ctx =
    let
        text14 =
            text 14

        left =
            text14 []

        middle =
            text14 [ Svg.Attributes.textAnchor "middle" ]

        right =
            text14 [ Svg.Attributes.textAnchor "end" ]

        ( x, y ) =
            ctx.node.label.pos

        neighbors =
            ctx.neighbors

        ancestors =
            ctx.ancestors

        outPipes =
            List.map (viewPipe ctx) neighbors

        outPorts =
            List.map left ctx.node.label.type'.outputs

        inPorts =
            List.map right ctx.node.label.type'.inputs

        flowDown i p =
            Svg.g [ Svg.Attributes.transform (translate 0 ((toFloat i) * portSpacing)) ] [ p ]

        translatedOutPorts =
            List.indexedMap flowDown outPorts

        translatedInPorts =
            List.indexedMap flowDown inPorts

        textNode =
            middle ctx.node.label.name

        box =
            Svg.rect
                [ Svg.Attributes.fill nodeColor
                , Svg.Attributes.width (px procWidth)
                , Svg.Attributes.height (px (procHeight ctx.node.label))
                , Svg.Attributes.rx (px 3)
                , Svg.Attributes.ry (px 3)
                , Svg.Attributes.style withDropShadow
                ]
                []
    in
        Svg.g
            [ Svg.Attributes.transform (translate x y)
            , Html.Events.onMouseDown (Click ctx.node.id)
            ]
            [ Svg.g [] outPipes
            , box
            , Svg.g [ Svg.Attributes.transform (translate (procWidth / 2) 0) ] [ textNode ]
            , Svg.g [ Svg.Attributes.transform (translate 0 0) ] translatedInPorts
            , Svg.g [ Svg.Attributes.transform (translate procWidth 0) ] translatedOutPorts
            ]


translate : Float -> Float -> String
translate x y =
    "translate(" ++ (toString x) ++ " " ++ (toString y) ++ ")"


px : number -> String
px n =
    (toString n) ++ "px"


nodeColor : String
nodeColor =
    "#fff"


text : Int -> List (Svg.Attribute Msg) -> String -> Svg.Svg Msg
text size attrs content =
    Svg.text'
        ([ Svg.Attributes.y (px size)
         , Svg.Attributes.fontSize (px size)
         ]
            ++ attrs
        )
        [ Svg.text content
        ]



-- SVG Extras


defs : List (Svg.Svg Msg)
defs =
    [ dropShadowFilter 3 1 1, arrowHead ]


arrowHead : Svg.Svg Msg
arrowHead =
    Svg.defs []
        [ Svg.marker
            [ Svg.Attributes.id "arrow-head"
            , Svg.Attributes.orient "auto"
            , Svg.Attributes.markerWidth "2"
            , Svg.Attributes.markerHeight "4"
            , Svg.Attributes.refX "0.1"
            , Svg.Attributes.refY "2"
            ]
            [ Svg.path [ Svg.Attributes.d "M0,0 V4 L2,2 Z", Svg.Attributes.fill "#000" ] [] ]
        ]


withDropShadow : String
withDropShadow =
    "filter:url(#dropshadow)"


dropShadowFilter : Float -> Int -> Int -> Svg.Svg Msg
dropShadowFilter blur dx dy =
    Svg.filter [ Svg.Attributes.id "dropshadow", Svg.Attributes.height "130%" ]
        [ Svg.feGaussianBlur
            [ Svg.Attributes.in' "SourceAlpha"
            , Svg.Attributes.stdDeviation (toString blur)
              -- blur
            ]
            []
        , Svg.feOffset
            [ Svg.Attributes.dx (toString dx)
              -- how much to offset
            , Svg.Attributes.dy (toString dy)
            , Svg.Attributes.result "offsetblur"
            ]
            []
        , Svg.feMerge []
            [ Svg.feMergeNode [] []
            , Svg.feMergeNode [ Svg.Attributes.in' "SourceGraphic" ] []
            ]
        ]
