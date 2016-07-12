module Main exposing (..)

import Html.App
import Task
import Html
import DataTypes exposing (..)
import Layout
import Window
import Graph
import Mouse
import Debug
import Platform.Sub


main : Program Never
main =
    Html.App.program
        { init = init
        , view = view
        , update = update
        , subscriptions = initialSubs
        }


initialSubs : Model -> Sub Msg
initialSubs model =
    Platform.Sub.batch [ windowResize, mouseMove ]


init : ( Model, Cmd Msg )
init =
    ( state0, Task.perform Resize Resize Window.size )


view : Model -> Html.Html Msg
view =
    Layout.layout


noFx : Model -> ( Model, Cmd Msg )
noFx m =
    ( m, Cmd.none )


update : Msg -> Model -> ( Model, Cmd Msg )
update msg model =
    case msg of
        Resize dim ->
            noFx <| { model | dim = dim }

        Click id ->
            noFx <| { model | selected = Just id }

        Unclick ->
            noFx <| { model | selected = Nothing }

        Drag p ->
            case model.selected of
                Nothing ->
                    noFx <| model

                Just id ->
                    noFx <| moveNode model id p


moveNode : Model -> Graph.NodeID -> Mouse.Position -> Model
moveNode model id pos =
    case Graph.nodeByID model.graph id of
        Nothing ->
            Debug.crash "undefined node! wtf!"

        Just oldNode ->
            let
                oldProc =
                    oldNode.label

                newNode =
                    Graph.Node id { oldProc | pos = ( Basics.toFloat pos.x, Basics.toFloat pos.y ) }
            in
                { model | graph = Graph.insertNode model.graph newNode }


windowResize : Sub Msg
windowResize =
    Window.resizes Resize


mouseMove : Sub Msg
mouseMove =
    Mouse.moves (translateMouse >> Drag)


translateMouse : Mouse.Position -> Mouse.Position
translateMouse p =
    { x = p.x - 300, y = p.y - 50 }
