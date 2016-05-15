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
    [ ( 0, 1, Pipe 1 0 )
    , ( 0, 2, Pipe 0 0 )
    , ( 1, 3, Pipe 0 0 )
    , ( 2, 3, Pipe 0 1 )
    ]


labeledProcs : List (Graph.Node Process)
labeledProcs =
    List.indexedMap Graph.Node procs


procs =
    [ { name = "something.MightFail"
      , pos = ( 50, 100 )
      , inputs = [ "INPUT" ]
      , outputs = [ "SUCCESS", "ERROR" ]
      }
    , { name = "log.Error"
      , pos = ( 300, 200 )
      , inputs = [ "INPUT" ]
      , outputs = [ "OUTPUT" ]
      }
    , { name = "fmt.Println"
      , pos = ( 400, 100 )
      , inputs = [ "INPUT" ]
      , outputs = [ "OUTPUT" ]
      }
    , { name = "base.Aggregator"
      , pos = ( 600, 200 )
      , inputs = [ "ACCUM", "NEXT" ]
      , outputs = [ "ACCUM" ]
      }
    ]


defaultProc =
    { name = "", pos = ( 0, 0 ), inputs = [], outputs = [] }


type alias Port =
    String


type alias PortID =
    Int


type alias Process =
    { name : String
    , pos : ( Float, Float )
    , inputs : List Port
    , outputs : List Port
    }


type alias Pipe =
    { output : PortID
    , input : PortID
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
