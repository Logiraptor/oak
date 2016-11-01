# Oak

[![Build Status](https://travis-ci.org/Logiraptor/oak.svg?branch=master)](https://travis-ci.org/Logiraptor/oak)
[![Go Report Card](https://goreportcard.com/badge/github.com/Logiraptor/oak)](https://goreportcard.com/report/github.com/Logiraptor/oak)
[![Coverage Status](https://coveralls.io/repos/github/Logiraptor/oak/badge.svg?branch=master)](https://coveralls.io/github/Logiraptor/oak?branch=master)

Oak is a flow based programming toolkit.

This repo is very immature and unstable. Issues and Contributions are welcome.

## Goals

This project is kind of an experiment for me. At this point it works for very simple programs.
I don't want to add features just because they're cool sounding, though.
My plan is to use this as a platform for all my side projects for a while.
I'll add features here only when I can't find a simple way to implement a
feature with the existing set of features or I find an annoying use case.

That said, the implementation goals as of now are:

- Dataflow / Flow Based Programming
- Scalable Processes within the dataflow graph
- Graphical UI for building the programs
	- I'd like to reuse existing projects if possible for this.
- Access to any Go package
- Multiple Targetable Environments (App Engine, CLI, WebServer, etc)

The target audience for this tool is myself at the moment. If this ends up being useful to other developers, I'm happy to share.

Next on the priority list:

- Test Coverage
- Better introspection of process types
	- I want to be able to name edges in the graph and 'split' data among components.
	- This should still be completely static analysis.
