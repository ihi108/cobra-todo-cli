#!/usr/bin/bash
go build .
if [ -f "cobra-todo-cli" ]; then
	mv cobra-todo-cli todo-cli
fi
