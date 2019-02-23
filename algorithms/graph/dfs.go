package graph

type dfsData struct {
	parent int
	start  int
	end    int
}

func DFS(g Graph, s int, f func(Node)) {
	var time int

	data := make(map[int]*dfsData)

	data[s] = &dfsData{
		parent: -1,
	}
	dfsVisit(g, s, f, data, &time)

	for i := range g.Nodes {
		if _, ok := data[i]; !ok {
			data[i] = &dfsData{
				parent: -1,
			}
			dfsVisit(g, i, f, data, &time)
		}
	}
}

func dfsVisit(g Graph, s int, f func(Node), data map[int]*dfsData, time *int) {
	*time++
	data[s].start = *time
	defer func() {
		f(g.Nodes[s])
		*time++
		data[s].end = *time
	}()

	for i, e := range g.Edges[s] {
		if _, ok := data[e.Target]; !ok {
			data[e.Target] = &dfsData{
				parent: s,
			}
			dfsVisit(g, e.Target, f, data, time)
			g.Edges[s][i].Type = TreeEdge
			continue
		}

		if data[e.Target].end == 0 {
			g.Edges[s][i].Type = BackwardEdge
			continue
		}

		if data[e.Target].start > data[s].start {
			g.Edges[s][i].Type = ForwardEdge
			continue
		}

		g.Edges[s][i].Type = CrossEdge
	}
}
