package main

type AlfredFunc interface {
	execute(str string) (string, error)
	tip() string
}
