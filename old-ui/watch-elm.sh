#!/bin/bash

cd elm && watch -tc elm make Main.elm --output ../public/assets/elm.js
