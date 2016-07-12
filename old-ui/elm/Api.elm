module Api exposing (..)

import Json.Decode exposing ((:=))
import Task
import Http


put : Json.Decode.Decoder a -> String -> Http.Body -> Task.Task Http.Error a
put decoder url body =
    let
        req =
            { verb = "PUT", body = body, url = url, headers = [ ( "Content-Type", "application/json" ) ] }
    in
        Http.send Http.defaultSettings req
            |> Http.fromJson decoder


delete : Json.Decode.Decoder a -> String -> Task.Task Http.Error a
delete decoder url =
    let
        req =
            { verb = "DELETE", body = Http.empty, url = url, headers = [] }
    in
        Http.send Http.defaultSettings req
            |> Http.fromJson decoder


decodeID : Json.Decode.Decoder { id : Int }
decodeID =
    Json.Decode.object1 (\id -> { id = id }) ("id" := Json.Decode.int)
