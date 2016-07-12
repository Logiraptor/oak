module Graph
    exposing
        ( Graph
        , NodeID
        , Node
        , Edge
        , NodeContext
        , fromNodesAndEdges
        , nodes
        , edges
        , mapNodes
        , outgoing
        , incoming
        , insertNode
        , nodeByID
        )

import Debug


type alias NodeID =
    Int


type alias Node n =
    { id : NodeID, label : n }


type alias Edge e =
    { from : NodeID, to : NodeID, label : e }


type Graph n e
    = GraphI (GraphInner n e)


type alias GraphInner n e =
    { nodes : List (Node n)
    , edges : List (Edge e)
    }


type alias NodeContext n e =
    { node : Node n
    , neighbors : List ( Edge e, Node n )
    , ancestors : List ( Edge e, Node n )
    }


gToStruct : Graph n e -> GraphInner n e
gToStruct g =
    case g of
        GraphI g ->
            g


mapNodes : (NodeContext n e -> a) -> Graph n e -> List a
mapNodes f g =
    List.map (f << buildContext g) (nodes g)


buildContext : Graph n e -> Node n -> NodeContext n e
buildContext g n =
    let
        outgoingEdges =
            List.filter (\e -> e.from == n.id) (edges g)

        incomingEdges =
            List.filter (\e -> e.to == n.id) (edges g)

        neighbors =
            List.map (\e -> ( e, nodeByID g e.to )) outgoingEdges

        ancestors =
            List.map (\e -> ( e, nodeByID g e.from )) incomingEdges

        keepExisting ( e, n ) =
            case n of
                Nothing ->
                    Nothing

                Just n ->
                    Just ( e, n )
    in
        { node = n
        , neighbors = List.filterMap keepExisting neighbors
        , ancestors = List.filterMap keepExisting ancestors
        }


insertNode : Graph n e -> Node n -> Graph n e
insertNode graph node =
    let
        otherNodes =
            List.filter (.id >> (/=) node.id) (nodes graph)

        struct =
            (gToStruct graph)
    in
        GraphI { struct | nodes = node :: otherNodes }


nodes : Graph n e -> List (Node n)
nodes g =
    (gToStruct g).nodes


edges : Graph n e -> List (Edge e)
edges g =
    (gToStruct g).edges


outgoing : Graph n e -> NodeID -> List ( e, n )
outgoing graph id =
    let
        outgoingEdges =
            List.filter (\e -> e.from == id) (edges graph)

        outgoingNodes =
            List.map (\e -> ( e.label, (unsafeGetNode graph e.to).label )) outgoingEdges
    in
        outgoingNodes


incoming : Graph n e -> NodeID -> List ( e, n )
incoming graph id =
    let
        incomingEdges =
            List.filter (\e -> e.to == id) (edges graph)

        incomingNodes =
            List.map (\e -> ( e.label, (unsafeGetNode graph e.from).label )) incomingEdges
    in
        incomingNodes


fromNodesAndEdges : List (Node n) -> List (Edge e) -> Graph n e
fromNodesAndEdges nodes edges =
    GraphI { nodes = nodes, edges = edges }


nodeByID : Graph n e -> NodeID -> Maybe (Node n)
nodeByID graph id =
    let
        available =
            nodes graph

        filtered =
            List.filter (.id >> (==) id) available
    in
        List.head filtered


unsafeGetNode : Graph n e -> NodeID -> Node n
unsafeGetNode g id =
    let
        result =
            nodeByID g id
    in
        case result of
            Maybe.Nothing ->
                Debug.crash "Unreachable code: Unknown graph node"

            Maybe.Just n ->
                n
