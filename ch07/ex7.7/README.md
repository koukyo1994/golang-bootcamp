# 練習問題7.7

20.0のデフォルト値は°を含んでいないのに、ヘルプメッセージが°を含んでいる理由を説明しなさい。

## 回答

`flag`パッケージではヘルプメッセージを表示するときの挙動を`flag.PrintDefaults()`メソッドで定義している。この定義は以下のようになっている。

```go
func (f *FlagSet) PrintDefaults() {
    f.VisitAll(func(flag *Flag) {
        var b strings.Builder
        fmt.Fprintf(&b, "  -%s", flag.Name) // Two spaces before -; see next two comments.
        name, usage := UnquoteUsage(flag)
        if len(name) > 0 {
            b.WriteString(" ")
            b.WriteString(name)
        }
        // Boolean flags of one ASCII letter are so common we
        // treat them specially, putting their usage on the same line.
        if b.Len() <= 4 { // space, space, '-', 'x'.
            b.WriteString("\t")
        } else {
            // Four spaces before the tab triggers good alignment
            // for both 4- and 8-space tab stops.
            b.WriteString("\n    \t")
        }
        b.WriteString(strings.ReplaceAll(usage, "\n", "\n    \t"))

        if !isZeroValue(flag, flag.DefValue) {
            if _, ok := flag.Value.(*stringValue); ok {
                // put quotes on the value
                fmt.Fprintf(&b, " (default %q)", flag.DefValue)
            } else {
                fmt.Fprintf(&b, " (default %v)", flag.DefValue)
            }
        }
        fmt.Fprint(f.Output(), b.String(), "\n")
    })
}
```

この中でデフォルト値を記述している部分は後半の`if !isZeroValue(flag, flag.DefValue)`以降であるが、ここでは、`flag.Value`を`*stringValue`にキャストできるかどうかで分岐している。キャストができない場合の挙動として、`%v`ヴァーヴを用いて値を表示することになっている。

今、`tempconv.CelsiusFlag`では`tempconv.Celsius`の値を設定するようになっているが、`tempconv.Celsius`に関して`%v`ヴァーヴを用いて表示をした場合に、`tempconv.Celsius`には`String()`メソッドが定義されているため、それを用いて値を表示する。`tempconv.Celsius.String()`の中では`fmt.Sprintf("%g°C", c)`によって値を表示するため、°を含んだ形でヘルプメッセージが表示される。
