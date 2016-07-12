module Messages exposing (..)


type Msg
    = Error Http.Error
    | ListComponents (List Component)
    | UpdateComponent Component
    | DeleteComponent Int
    | Finished
