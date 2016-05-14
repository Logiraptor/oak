module Layout exposing (..)

-- where

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
        procs =
            Graph.nodes model.graph

        procElems =
            List.map (viewNode model) procs

        pipes =
            Graph.edges model.graph

        pipeElems =
            List.map (viewPipe model) pipes
    in
        Html.div [ Html.Attributes.class "editor" ]
            [ Svg.svg
                [ Svg.Attributes.width ("100%")
                , Svg.Attributes.height ("100%")
                ]
                [ Svg.g [] pipeElems
                , Svg.g [] procElems
                , Svg.g [] filters
                , Svg.g []
                    [ text 14 [] (toString model)
                    ]
                ]
            ]


viewPipe : Model -> Graph.Edge Pipe -> Svg.Svg Msg
viewPipe model pipe =
    let
        from =
            Graph.nodeByID model.graph pipe.from

        to =
            Graph.nodeByID model.graph pipe.to
    in
        case ( from, to ) of
            ( Just from, Just to ) ->
                line ( procWidth + (fst from.label.pos), snd from.label.pos )
                    to.label.pos
                    [ Svg.Attributes.stroke "#000"
                    , Svg.Attributes.fill "none"
                    ]
                    []

            _ ->
                Svg.text ""


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
                , pre "C " ((x2 + x1) / 2) y1
                , pre "" ((x2 + x1) / 2) y2
                , pre "" x2 y2
                ]
    in
        Svg.path << (::) (Svg.Attributes.d d)


viewNode : Model -> Graph.Node Process -> Svg.Svg Msg
viewNode model node =
    let
        text14 =
            text 14

        left =
            text14 []

        right =
            text14 [ Svg.Attributes.textAnchor "end" ]

        ( x, y ) =
            node.label.pos

        outgoing =
            Graph.outgoing model.graph node.id

        incoming =
            Graph.incoming model.graph node.id

        inPorts =
            List.map (fst >> .input >> right) incoming

        outPorts =
            List.map (fst >> .output >> left) outgoing

        flowDown i p =
            Svg.g [ Svg.Attributes.transform (translate 0 ((toFloat i) * 25)) ] [ p ]

        translatedInPorts =
            List.indexedMap flowDown inPorts

        translatedOutPorts =
            List.indexedMap flowDown outPorts

        textNode =
            left node.label.name

        box =
            Svg.rect
                [ Svg.Attributes.fill nodeColor
                , Svg.Attributes.width (px procWidth)
                , Svg.Attributes.height (px 50)
                , Svg.Attributes.rx (px 3)
                , Svg.Attributes.ry (px 3)
                , Svg.Attributes.style withDropShadow
                ]
                []
    in
        Svg.g
            [ Svg.Attributes.transform (translate x y)
            , Html.Events.onMouseDown (Click node.id)
            ]
            [ box
            , textNode
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



-- Filters


filters : List (Svg.Svg Msg)
filters =
    [ dropShadowFilter 3 1 1 ]


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
