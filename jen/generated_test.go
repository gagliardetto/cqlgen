package jen_test

import (
	"testing"

	. "github.com/gagliardetto/cqlgen/jen"
)

var gencases = []tc{
	{
		desc: `bool group`,
		code: BlockFunc(func(g *Group) {
			g.Int()
		}),
		expect: `{
		int
		}`,
	},
	{
		desc: `paramsfunc group`,
		// Don't do this! ListFunc used to kludge Group.ParamsFunc usage without
		// syntax error.
		code:   Id("a").ListFunc(func(lg *Group) { lg.ParamsFunc(func(cg *Group) { cg.Lit(1) }) }),
		expect: `a(1)`,
	},
	{
		desc:   `paramsfunc func`,
		code:   Id("a").Add(ParamsFunc(func(g *Group) { g.Lit(1) })),
		expect: `a(1)`,
	},
	{
		desc:   `paramsfunc statement`,
		code:   Id("a").ParamsFunc(func(g *Group) { g.Lit(1) }),
		expect: `a(1)`,
	},
	{
		desc: `params group`,
		// Don't do this! ListFunc used to kludge Group.Params usage without
		// syntax error.
		code:   Id("a").ListFunc(func(g *Group) { g.Params(Lit(1)) }),
		expect: `a(1)`,
	},
	{
		desc:   `params func`,
		code:   Id("a").Add(Params(Lit(1))),
		expect: `a(1)`,
	},
	{
		desc: `callfunc group`,
		// Don't do this! ListFunc used to kludge Group.CallFunc usage without
		// syntax error.
		code:   Id("a").ListFunc(func(lg *Group) { lg.CallFunc(func(cg *Group) { cg.Lit(1) }) }),
		expect: `a(1)`,
	},
	{
		desc:   `callfunc func`,
		code:   Id("a").Add(CallFunc(func(g *Group) { g.Lit(1) })),
		expect: `a(1)`,
	},
	{
		desc: `call group`,
		// Don't do this! ListFunc used to kludge Group.Call usage without
		// syntax error.
		code:   Id("a").ListFunc(func(g *Group) { g.Call(Lit(1)) }),
		expect: `a(1)`,
	},
	{
		desc:   `call func`,
		code:   Id("a").Add(Call(Lit(1))),
		expect: `a(1)`,
	},
	{
		desc:   `blockfunc group`,
		code:   BlockFunc(func(g *Group) { g.BlockFunc(func(g *Group) {}) }),
		expect: `{{}}`,
	},
	{
		desc:   `block group`,
		code:   BlockFunc(func(g *Group) { g.Block() }),
		expect: `{{}}`,
	},
	{
		desc:   `indexfunc statement`,
		code:   Id("a").IndexFunc(func(g *Group) { g.Lit(1) }),
		expect: `a[1]`,
	},
	{
		desc:   `indexfunc func`,
		code:   Id("a").Add(IndexFunc(func(g *Group) { g.Lit(1) })),
		expect: `a[1]`,
	},
	{
		desc:   `index func`,
		code:   Id("a").Add(Index(Lit(1))),
		expect: `a[1]`,
	},
	{
		desc: `listfunc statement`,
		code: Add(Null()).ListFunc(func(lg *Group) {
			lg.Id("a")
			lg.Id("b")
		}).Op("=").Id("c"),
		expect: `a, b = c`,
	},
	{
		desc: `listfunc func`,
		code: ListFunc(func(lg *Group) {
			lg.Id("a")
			lg.Id("b")
		}).Op("=").Id("c"),
		expect: `a, b = c`,
	},
	{
		desc: `listfunc group`,
		code: BlockFunc(func(bg *Group) {
			bg.ListFunc(func(lg *Group) {
				lg.Id("a")
				lg.Id("b")
			}).Op("=").Id("c")
		}),
		expect: `{
		a, b = c
		}`,
	},
	{
		desc: `list group`,
		code: BlockFunc(func(g *Group) { g.List(Id("a"), Id("b")).Op("=").Id("c") }),
		expect: `{
		a, b = c
		}`,
	},
	{
		desc:   `parens func`,
		code:   Parens(Lit(1)),
		expect: `(1)`,
	},
	{
		desc: `parens group`,
		code: BlockFunc(func(g *Group) { g.Parens(Lit(1)) }),
		expect: `{
		(1)
		}`,
	},
}

func TestGen(t *testing.T) {
	caseTester(t, gencases)
}
