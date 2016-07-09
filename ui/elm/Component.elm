module Component exposing (..)

import Task
import Http
import Json.Encode
import Json.Decode exposing ((:=))
import Api exposing (..)


type alias Component =
    { name : String, id : Int }


decodeComponent : Json.Decode.Decoder Component
decodeComponent =
    Json.Decode.object2 Component
        (Json.Decode.oneOf [ "name" := Json.Decode.string, Json.Decode.succeed "" ])
        ("id" := Json.Decode.int)


encodeComponent : Component -> Json.Encode.Value
encodeComponent c =
    Json.Encode.object
        [ ( "id", Json.Encode.int c.id )
        , ( "name", Json.Encode.string c.name )
        ]


createComponent : Component -> Cmd
createComponent component =
    encodeComponent component
        |> Json.Encode.encode 0
        |> Http.string
        |> Http.post decodeID "http://localhost:3000/components/"
        |> Task.perform Error finish


updateComponent : Component -> Cmd
updateComponent component =
    encodeComponent component
        |> Json.Encode.encode 0
        |> Http.string
        |> put decodeID ("http://localhost:3000/components/" ++ (toString component.id))
        |> Task.perform Error finish


deleteComponent : Int -> Cmd
deleteComponent id =
    delete (Json.Decode.succeed id) ("http://localhost:3000/components/" ++ (toString id))
        |> Task.perform Error finish


listComponents : Cmd
listComponents =
    Http.get decodeComponents "http://localhost:3000/components"
        |> Task.perform Error ListComponents
