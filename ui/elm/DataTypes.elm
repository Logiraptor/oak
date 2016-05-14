module DataTypes exposing (..)

--where

import Graph
import Mouse


state0 : Model
state0 =
    { graph =
        Graph.fromNodesAndEdges labeledProcs
            labeledPipes
    , dim = { width = 0, height = 0 }
    , selected = Nothing
    }


labeledPipes : List (Graph.Edge Pipe)
labeledPipes =
    List.map (\( from, to, p ) -> Graph.Edge from to p) pipes


pipes =
    [ ( 0, 1, { defaultPipe | output = "SUCCESS" } )
    , ( 0, 2, { defaultPipe | output = "ERROR" } )
    ]


labeledProcs : List (Graph.Node Process)
labeledProcs =
    List.indexedMap Graph.Node procs


procs =
    [ { defaultProc | name = "strings.ToUpper", pos = ( 0, 100 ) }
    , { defaultProc | name = "fmt.Println", pos = ( 300, 200 ) }
    , { defaultProc | name = "fmt.Scanln", pos = ( 400, 100 ) }
    ]


defaultProc =
    { name = "", pos = ( 0, 0 ) }


defaultPipe =
    { input = "INPUT", output = "OUTPUT" }


type alias Process =
    { name : String
    , pos : ( Float, Float )
    }


type alias Pipe =
    { input : String
    , output : String
    }


type alias FlowGraph =
    Graph.Graph Process Pipe


type alias Model =
    { graph : FlowGraph
    , dim : { width : Int, height : Int }
    , selected : Maybe Graph.NodeID
    }


type Msg
    = Resize { width : Int, height : Int }
    | Click Graph.NodeID
    | Unclick
    | Drag Mouse.Position
