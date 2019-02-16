package str

import (
	"fmt"
	"log"
)

type Operation struct {
	Type OperationType
	A    byte
	B    byte
}

func (o Operation) String() string {
	switch o.Type {
	case Copy, Delete, Insert:
		return fmt.Sprintf("%s %s", o.Type, string(o.A))
	case Replace:
		return fmt.Sprintf("%s %s by %s", o.Type, string(o.A), string(o.B))
	}
	return ""
}

//go:generate stringer -type OperationType

type OperationType int

const (
	NoOp OperationType = iota
	Copy
	Replace
	Delete
	Insert
)

var operationCosts = map[OperationType]int{
	Copy:    -1,
	Replace: 1,
	Delete:  2,
	Insert:  2,
}

func Transform(a, b string) []Operation {
	_, ops := computeTables(a, b)
	return assembleTransformation(ops, len(a), len(b))
}

func assembleTransformation(ops [][]Operation, i, j int) []Operation {
	if i == 0 && j == 0 {
		return nil
	}

	op := ops[i][j]
	switch op.Type {
	case Replace, Copy:
		return append(assembleTransformation(ops, i-1, j-1), ops[i][j])
	case Delete:
		return append(assembleTransformation(ops, i-1, j), ops[i][j])
	case Insert:
		return append(assembleTransformation(ops, i, j-1), ops[i][j])
	}
	log.Panicf("Invalid operation type: %s (%#v)", op.Type, op.Type)
	return nil
}

func computeTables(a, b string) ([][]int, [][]Operation) {
	costs := make([][]int, len(a)+1)
	ops := make([][]Operation, len(a)+1)
	for i := range costs {
		costs[i] = make([]int, len(b)+1)
		ops[i] = make([]Operation, len(b)+1)

		if i == 0 {
			for j := range costs[i] {
				if j == 0 {
					continue
				}

				costs[i][j] = j * operationCosts[Insert]
				ops[i][j] = Operation{
					Type: Insert,
					A:    b[j-1],
				}
			}
			continue
		}

		costs[i][0] = i * operationCosts[Delete]
		ops[i][0] = Operation{
			Type: Delete,
			A:    (a[i-1]),
		}
	}

	for i := 1; i < len(costs); i++ {
		for j := 1; j < len(costs[i]); j++ {
			if a[i-1] == b[j-1] {
				costs[i][j] = costs[i-1][j-1] + operationCosts[Copy]
				ops[i][j] = Operation{
					Type: Copy,
					A:    a[i-1],
				}
			} else {
				costs[i][j] = costs[i-1][j-1] + operationCosts[Replace]
				ops[i][j] = Operation{
					Type: Replace,
					A:    a[i-1],
					B:    b[j-1],
				}
			}

			if costs[i][j] > costs[i][j-1]+operationCosts[Insert] {
				costs[i][j] = costs[i][j-1] + operationCosts[Insert]
				ops[i][j] = Operation{
					Type: Insert,
					A:    b[j-1],
				}
			}
			if costs[i][j] > costs[i-1][j]+operationCosts[Delete] {
				costs[i][j] = costs[i-1][j] + operationCosts[Delete]
				ops[i][j] = Operation{
					Type: Delete,
					A:    a[i-1],
				}
			}
		}
	}

	return costs, ops
}
