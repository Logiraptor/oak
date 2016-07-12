module Application where


type alias Edge =
    { from: (Int, Int)
    , to : (Int, Int)
    }


type alias Node =
    { process: Process
    , position: (Float, Float)
    }

type alias Graph =
    { nodes: List Node
    , edges: List Edge
    }


type alias Port =
    { typ: String
    }


type alias Process =
    { name: String
    , inputs: List Port
    , outputs: List Port
    }
