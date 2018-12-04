package evaluator

import (
	"Simplex-Simia/object"
	"fmt"
	"strconv"
)

var builtins = map[string]*object.Builtin{
	"len": &object.Builtin{Fn: func(args ...object.Object) object.Object {
		if len(args) != 1 {
			return newError("wrong number of arguments. got=%d, want=1",
				len(args))
		}

		switch arg := args[0].(type) {
		case *object.Array:
			return &object.Integer{Value: int64(len(arg.Elements))}
		case *object.String:
			return &object.Integer{Value: int64(len(arg.Value))}
		// case *object.Integer:
		// 	return &object.Integer{Value: int64(len(strconv.Itoa(int(arg.Value))))}
		default:
			return newError("argument to `len` not supported, got %s",
				args[0].Type())
		}
	},
	},
	"puts": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
	"first": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `first` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			if len(arr.Elements) > 0 {
				return arr.Elements[0]
			}

			return NULL
		},
	},
	"last": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `last` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				return arr.Elements[length-1]
			}

			return NULL
		},
	},
	"rest": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `rest` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)
			if length > 0 {
				newElements := make([]object.Object, length-1, length-1)
				copy(newElements, arr.Elements[1:length])
				return &object.Array{Elements: newElements}
			}

			return NULL
		},
	},
	"push": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("wrong number of arguments. got=%d, want=2",
					len(args))
			}
			if args[0].Type() != object.ARRAY_OBJ {
				return newError("argument to `push` must be ARRAY, got %s",
					args[0].Type())
			}

			arr := args[0].(*object.Array)
			length := len(arr.Elements)

			newElements := make([]object.Object, length+1, length+1)
			copy(newElements, arr.Elements)
			newElements[length] = args[1]

			return &object.Array{Elements: newElements}
		},
	},
	"print": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var blank object.Object
			for _, arg := range args {
				fmt.Print(arg.Inspect())
			}
			return blank
		},
	},
	"println": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			var blank object.Object
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return blank
		},
	},
	"c": &object.Builtin{Fn: func(args ...object.Object) object.Object {
		var blank object.Object
		if len(args) != 2 {
			return newError("wrong number of arguments. got=%d, want=2",
				len(args))
		}

		if args[0].Type() != object.INTEGER_OBJ {
			return newError("argument to `c` not supported, got %s",
				args[0].Type())
		} else if args[1].Type() != object.INTEGER_OBJ {
			return newError("argument to `c` not supported, got %s",
				args[1].Type())
		}

		val1, err1 := strconv.Atoi(args[0].Inspect())
		val2, err2 := strconv.Atoi(args[1].Inspect())
		if err1 != nil || err2 != nil || val1 < val2 {
			return newError("val2=%d must not be greater than val1=%d", val2, val1)
		}
		fmt.Println(coefficient(val1, val2))
		return blank
	},
	},
}

func coefficient(n, k int) int64 {
	if k == 0 || k == n {
		return 1
	}

	return coefficient(n-1, k-1) + coefficient(n-1, k)
}
