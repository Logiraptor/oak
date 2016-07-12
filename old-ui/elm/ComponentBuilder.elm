module ComponentBuilder exposing (..)

import Html.App
import Html
import Html.Attributes
import Html.Events
import Platform.Cmd
import Platform.Sub
import Http
import Json.Decode exposing ((:=))
import Json.Encode
import Task
import Component exposing (..)


main =
    Html.App.program
        { init = init
        , view = view
        , update = update
        , subscriptions = \_ -> Platform.Sub.none
        }


init : ( Model, Cmd )
init =
    ( state0, initialCmds )


initialCmds : Cmd
initialCmds =
    let
        t =
            Http.get decodeComponents "http://localhost:3000/components"
    in
        Task.perform Error ListComponents t


decodeComponents : Json.Decode.Decoder (List Component)
decodeComponents =
    Json.Decode.list decodeComponent


view : Model -> Html.Html Msg
view model =
    Html.div []
        [ Html.div [] (List.map viewComponent model.components)
        , Html.div [] [ Html.button [ Html.Events.onClick (UpdateComponent { id = newID model.components, name = "" }) ] [ Html.text "Add Component" ] ]
        , Html.div [] [ Html.text (Basics.toString model.err) ]
        ]


viewComponent : Component -> Html.Html Msg
viewComponent component =
    Html.div []
        [ Html.input [ Html.Attributes.value component.name, Html.Events.onInput (\name -> (UpdateComponent { component | name = name })) ] []
        , Html.button [ Html.Events.onClick (DeleteComponent component.id) ] [ Html.text "Delete" ]
        ]


update : Msg -> Model -> ( Model, Cmd )
update msg model =
    let
        noFx m =
            ( m, Platform.Cmd.none )
    in
        case msg of
            Error e ->
                { model | err = Just e } |> noFx

            ListComponents r ->
                { model | components = r } |> noFx

            UpdateComponent c ->
                let
                    ( components, cmd ) =
                        replaceComponent model.components c
                in
                    ( { model | components = components }, cmd )

            DeleteComponent id ->
                let
                    ( components, cmd ) =
                        removeComponent model.components id
                in
                    ( { model | components = components }, cmd )

            Finished ->
                model |> noFx


removeComponent : List Component -> Int -> ( List Component, Cmd )
removeComponent components id =
    case components of
        [] ->
            ( components, Platform.Cmd.none )

        h :: t ->
            if h.id == id then
                ( t, deleteComponent id )
            else
                let
                    ( components, cmd ) =
                        removeComponent t id
                in
                    ( h :: components, cmd )


replaceComponent : List Component -> Component -> ( List Component, Cmd )
replaceComponent components component =
    case components of
        [] ->
            ( component :: components, createComponent component )

        h :: t ->
            if h.id == component.id then
                ( component :: t, updateComponent component )
            else
                let
                    ( components, cmd ) =
                        replaceComponent t component
                in
                    ( h :: components, cmd )


newID : List { a | id : Int } -> Int
newID list =
    case List.maximum (List.map .id list) of
        Maybe.Nothing ->
            0

        Maybe.Just n ->
            n + 1


state0 : Model
state0 =
    { components = [], err = Nothing }


type alias Model =
    { components : List Component, err : Maybe Http.Error }


type alias Record r =
    { r | id : Int }


type alias Cmd =
    Platform.Cmd.Cmd Msg



--- HTTP


finish : a -> Msg
finish _ =
    Finished
