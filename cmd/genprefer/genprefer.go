package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"go/format"
	"io"
	"os"
	"os/signal"
	"strings"
)

func (r *Rule) setDefaults() {
	cleanArgs := func(a []string) string {
		clean := make([]string, len(a))
		for i := range a {
			clean[i] = strings.ReplaceAll(a[i], "$", "")
		}
		return args(clean)
	}

	if r.Before == "" {
		r.Before = "require." + r.Match[0].Name + "(t, " + cleanArgs(r.Match[0].Args) + ")"
	}
	if r.After == "" {
		r.After = "require." + r.Suggest.Name + "(t, " + cleanArgs(r.Suggest.Args) + ")"
	}
	if len(r.Tags) == 0 {
		r.Tags = []string{"diagnostic"}
	}

	if r.AssocMatch {
		for _, m := range r.Match {
			m.Args = append([]string(nil), m.Args...)
			for left, right := 0, len(m.Args)-1; left < right; left, right = left+1, right-1 {
				m.Args[left], m.Args[right] = m.Args[right], m.Args[left]
			}
			r.Match = append(r.Match, m)
		}
	}
}

func rules() (r []Rule) {
	const (
		condIsError = `m["a"].Type.Is("error")`
	)

	r = []Rule{
		{
			Name:       "preferLen",
			Summary:    "Prefer require.Len instead of comparing length.",
			AssocMatch: true,
			Match: []Method{
				{Name: "Equal", Args: []string{"$length", "len($a)"}},
				{Name: "Exactly", Args: []string{"$length", "len($a)"}},
			},
			Suggest: Method{Name: "Len", Args: []string{"$a", "$length"}},
		},
		{
			Name:       "preferEmpty",
			Summary:    "Prefer require.Empty instead of comparing length.",
			AssocMatch: true,
			Match: []Method{
				{Name: "Equal", Args: []string{"0", "len($a)"}},
			},
			Suggest: Method{Name: "Empty", Args: []string{"$a"}},
		},
		{
			Name:       "preferNil",
			Summary:    "Prefer require.Nil instead of comparing to nil.",
			AssocMatch: true,
			Match: []Method{
				{Name: "Equal", Args: []string{"nil", "$a"}},
				{Name: "Exactly", Args: []string{"nil", "$a"}},
				{Name: "Same", Args: []string{"nil", "$a"}},
			},
			Suggest: Method{Name: "Nil", Args: []string{"$a"}},
		},
		{
			Name:       "preferNotNil",
			Summary:    "Prefer require.NotNil instead of comparing to nil.",
			AssocMatch: true,
			Match: []Method{
				{Name: "NotEqual", Args: []string{"nil", "$a"}},
			},
			Suggest: Method{Name: "NotNil", Args: []string{"$a"}},
		},
		{
			Name:    "preferErrorIs",
			Summary: "Prefer require.ErrorIs instead of comparing error values by Equal.",
			Match: []Method{
				{
					Name: "Equal",
					Args: []string{"$target", "$a"},
					Cond: condIsError + `&& m["target"].Object.Is("Var")`,
				},
			},
			Suggest: Method{Name: "ErrorIs", Args: []string{"$a", "$target"}},
		},
		{
			Name:    "preferErrorAs",
			Summary: "Prefer require.ErrorAs instead of comparing error values by Equal.",
			Match: []Method{
				{
					Name: "Equal",
					Args: []string{"$target", "$a"},
					Cond: condIsError + `&& !m["target"].Object.Is("Var")`,
				},
			},
			Suggest: Method{Name: "ErrorAs", Args: []string{"$a", "$target"}},
		},
		{
			Name:    "preferError",
			Summary: "Prefer require.NoError instead of comparing to nil.",
			Match: []Method{
				{Name: "NotNil", Args: []string{"$a"}, Cond: condIsError},
			},
			Suggest: Method{Name: "Error", Args: []string{"$a"}},
		},
		{
			Name:    "preferNoError",
			Summary: "Prefer require.NoError instead of comparing to nil.",
			Match: []Method{
				{Name: "Nil", Args: []string{"$a"}, Cond: condIsError},
			},
			Suggest: Method{Name: "NoError", Args: []string{"$a"}},
		},
	}

	for i := range r {
		r[i].setDefaults()
	}

	return r
}

func generate(w io.Writer) error {
	buf := bytes.Buffer{}

	if err := tmpl.ExecuteTemplate(&buf, "main", Config{
		Rules: rules(),
	}); err != nil {
		return err
	}

	formatted, err := format.Source(buf.Bytes())
	if err != nil {
		return err
	}

	if _, err := w.Write(formatted); err != nil {
		return err
	}

	return nil
}

func run(ctx context.Context) error {
	output := flag.String("output", "", "output file")
	flag.Parse()

	var w io.Writer = os.Stdout
	if p := *output; p != "" {
		f, err := os.Create(p)
		if err != nil {
			return err
		}
		defer func() {
			_ = f.Close()
		}()
		w = f
	}

	return generate(w)
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	if err := run(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
