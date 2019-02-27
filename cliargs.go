package cliargs

import (
	"bufio"
	"os"
	"strings"
	"unicode"

	"github.com/thamaji/slices"
	cli "gopkg.in/urfave/cli.v2"
)

type args []string

// DefaultFuncWhenSingleHyphen は，cli.Appの引数の最後に "-" が与えられた時のデフォルトの挙動を示す．
// 標準入力から空白区切りの引数列を読み，stringのスライスにして返す．
// 空白文字の定義は， unicode.IsSpace() == true となる rune型 とする．
var DefaultFuncWhenSingleHyphen = func() ([]string, error) {
	scanner := bufio.NewScanner(os.Stdin)
	args := make([]string, 0, 16)
	separateWhiteSpaceFunc := func(r rune) bool {
		return unicode.IsSpace(r)
	}
	for scanner.Scan() {
		line := scanner.Text()
		lineArgs := strings.FieldsFunc(line, separateWhiteSpaceFunc)
		args = append(args, lineArgs...)
	}
	return args, nil
}

// WrapPOSIXLike は，argsに与えられたcli.Argsに "-" が与えられているならば， "-" を全て取り除いた上で新たなcli.Argsにして返す．
// cli.Args の末尾に "-" が与えられていた場合，"-"を取り除いた上で，onLastSingleHyphen()により新たな引数の入力を受け付ける．
func WrapPOSIXLike(oldArgs cli.Args, onLastSingleHyphen func() ([]string, error)) (cli.Args, error) {
	var newArgsStrings []string
	oldArgsSlice := oldArgs.Slice()

	// 引数列に "-" が入っているなら "-"を取り除く
	excluded := slices.FilterStringFunc(oldArgsSlice, func(s string) bool {
		return s != "-"
	})
	newArgsStrings = excluded
	// 引数列の末尾に "-" が入っていたならば，入力を受け付ける
	if oldArgsSlice[oldArgs.Len()-1] == "-" {
		sequence, err := onLastSingleHyphen()
		if err != nil {
			return nil, err
		}
		return expand(newArgsStrings, sequence)
	}

	newArgs := args(newArgsStrings)
	return &newArgs, nil
}

func expand(oldArgs []string, newArgs []string) (cli.Args, error) {
	expanded := make([]string, 0, len(oldArgs)+len(newArgs))
	expanded = append(expanded, oldArgs...)
	expanded = append(expanded, newArgs...)
	args := args(expanded)
	return cli.Args(&args), nil
}

func (a *args) Get(n int) string {
	if len(*a) > n {
		return (*a)[n]
	}
	return ""
}

func (a *args) First() string {
	return a.Get(0)
}

func (a *args) Tail() []string {
	if a.Len() >= 2 {
		tail := []string((*a)[1:])
		ret := make([]string, len(tail))
		copy(ret, tail)
		return ret
	}
	return []string{}
}

func (a *args) Len() int {
	return len(*a)
}

func (a *args) Present() bool {
	return a.Len() != 0
}

func (a *args) Slice() []string {
	ret := make([]string, len(*a))
	copy(ret, []string(*a))
	return ret
}
