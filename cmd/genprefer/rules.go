package main

type Method struct {
	Name string
	Args []string
	Cond string
}

type Rule struct {
	Name       string
	Summary    string
	Before     string
	After      string
	AssocMatch bool
	Tags       []string
	Match      []Method
	Suggest    Method
}

type Config struct {
	Rules []Rule
}
