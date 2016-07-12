module JsonToFunc (..) where

import Graphics.Element
import Json.Decode as Js exposing (Decoder)
import Dict exposing (Dict)
import Array exposing (Array)


main =
  Graphics.Element.show stuff


stuff =
  Js.decodeString (buildDecoder ()) input


type Data
  = Object (Dict String Data)
  | Array (Array Data)
  | String String
  | Int Int
  | Float Float
  | Bool Bool
  | Null


buildDecoder : () -> Decoder Data
buildDecoder _ =
  Js.oneOf
    [ buildObjectDecoder ()
    , buildArrayDecoder ()
    , stringDecoder
    , intDecoder
    , floatDecoder
    , boolDecoder
    , nullDecoder
    ]


buildObjectDecoder : () -> Decoder Data
buildObjectDecoder _ =
  Js.succeed () `Js.andThen` (buildDecoder >> Js.dict >> Js.map Object)


buildArrayDecoder : () -> Decoder Data
buildArrayDecoder _ =
  Js.succeed () `Js.andThen` (buildDecoder >> Js.array >> Js.map Array)


decoder : Decoder Data
decoder =
  buildDecoder ()


objectDecoder : Decoder Data
objectDecoder =
  buildObjectDecoder ()


arrayDecoder : Decoder Data
arrayDecoder =
  buildArrayDecoder ()


stringDecoder : Decoder Data
stringDecoder =
  Js.string |> Js.map String


intDecoder : Decoder Data
intDecoder =
  Js.int |> Js.map Int


floatDecoder : Decoder Data
floatDecoder =
  Js.float |> Js.map Float


boolDecoder : Decoder Data
boolDecoder =
  Js.bool |> Js.map Bool


nullDecoder : Decoder Data
nullDecoder =
  Js.null Null


input =
  """
  {"name": "Patrick"}
  """
