module SceneTree where

import Graphics.Collage

type SceneTree =
    Leaf (Float, Float) Graphics.Collage.Form
  | Move (Float, Float) SceneTree
  | Group (List SceneTree)
  | NoCollide SceneTree


toForm : SceneTree -> Graphics.Collage.Form
toForm tree =
    case tree of
        Leaf _ f ->
            f
        Move pos f ->
            Graphics.Collage.move pos (toForm f)
        Group nodes ->
            Graphics.Collage.group (List.map toForm nodes)
        NoCollide t ->
            toForm t


collide : (Float, Float) -> SceneTree -> Bool
collide (x, y) tree =
    case tree of
        Leaf (w, h) _ ->
            x < w && x > 0 && y < h && y > 0
        Move (lx, ly) f ->
            collide (x-lx, y-lx) f
        Group nodes ->
            List.any (collide (x, y)) nodes
        NoCollide t ->
            False
