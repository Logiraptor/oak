module Main exposing (..)

import Html
import Navigation
import UrlParser
import Platform.Cmd
import Platform.Sub


type Page
    = Index


type alias Model =
    Int


type Msg
    = Nothing


type alias Cmd =
    Platform.Cmd.Cmd Msg


type alias Sub =
    Platform.Sub.Sub Msg


main : Program Never
main =
    Navigation.program (Navigation.makeParser hashParser)
        { init = ( state0, Platform.Cmd.none )
        , update = update
        , view = view
        , subscriptions = subscriptions
        , urlUpdate = urlUpdate
        }


hashParser : Navigation.Location -> Result String Page
hashParser location =
    UrlParser.parse identity pageParser (String.dropLeft 1 location.hash)


pageParser : Parser (Page -> a) a
pageParser =
    UrlParser.oneOf
        [ UrlParser.format Index (UrlParser.s "home")
        ]


update : Msg -> Model -> ( Model, Cmd )
update msg model =
    ( model, Platform.Cmd.none )


subscriptions : Model -> Sub
subscriptions model =
    Platform.Sub.none


urlUpdate : Page -> Model -> ( Model, Cmd )
urlUpdate page model =
    ( model, Platform.Cmd.none )


view : Model -> Html.Html Msg
view model =
    Html.h1 [] []


state0 : Model
state0 =
    0
