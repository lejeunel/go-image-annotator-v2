package html

import (
	gp "maragu.dev/gomponents"
	gh "maragu.dev/gomponents/html"
)

type TableRow struct {
	Values []gp.Node
}

func (r TableRow) Render() gp.Node {
	return gh.Tr(gh.Class("text-neutral-1000"),
		gp.Map(r.Values, func(node gp.Node) gp.Node {
			return gh.Td(gh.Class("px-5 py-4 text-sm whitespace-nowrap"),
				node)
		}))

}

func MyTable(fields []string, rows []TableRow) gp.Node {
	return gh.Table(gh.Class("min-w-full divide-y divide-neutral-200"),
		TableHeader(fields),
		TableBody(rows),
	)
}

func TableHeader(fields []string) gp.Node {
	return gh.THead(gh.Tr(gh.Class("text-neutral-500"),
		gp.Map(fields, func(f string) gp.Node {
			return gh.Th(gh.Class("px-5 py-3 text-xs font-medium text-left uppercase"), gp.Text(f))
		})))
}

func TableBody(rows []TableRow) gp.Node {
	return gh.TBody(gh.Class("divide-y divide-neutral-200"),
		gp.Map(rows, func(r TableRow) gp.Node {
			return r.Render()

		}),
	)
}
