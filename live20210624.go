package live20210624

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/ast/inspector"
)

const doc = "live20210624 is ..."

// Analyzer is ...
var Analyzer = &analysis.Analyzer{
	Name: "live20210624",
	Doc:  doc,
	Run:  run,
	Requires: []*analysis.Analyzer{
		inspect.Analyzer,
	},
}

func run(pass *analysis.Pass) (interface{}, error) {
	inspect := pass.ResultOf[inspect.Analyzer].(*inspector.Inspector)

	nodeFilter := []ast.Node{
		(*ast.FuncDecl)(nil),
		(*ast.FuncLit)(nil),
		(*ast.Ident)(nil),
	}

	inspect.Preorder(nodeFilter, func(n ast.Node) {
		switch n := n.(type) {
		case *ast.FuncDecl:
			if n.Type.Params != nil {
				for _, arg := range n.Type.Params.List {
					chType, ok := arg.Type.(*ast.ChanType)
					if !ok {
						continue
					}
					if chType.Dir == ast.RECV|ast.SEND {
						pass.Reportf(n.Pos(), "channel argument should be directed")
					}
				}
			}
			if n.Type.Results != nil {
				for _, res := range n.Type.Results.List {
					chType, ok := res.Type.(*ast.ChanType)
					if !ok {
						continue
					}
					if chType.Dir == ast.RECV|ast.SEND {
						pass.Reportf(n.Pos(), "channel result should be directed")
					}
				}
			}
		case *ast.Ident:
			if n.Name == "gopher" {
				pass.Reportf(n.Pos(), "identifier is gopher")
			}
		}
	})

	return nil, nil
}
