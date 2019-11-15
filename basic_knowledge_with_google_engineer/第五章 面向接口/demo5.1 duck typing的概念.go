package main

/*
duck typing
"像鸭子走路，像鸭子叫(长得像鸭子)，那么就是鸭子"
描述事物的外部行为而非内部结构
严格说go属于结构化系统，类似duck typing

python中的duck typing
def download(retriever):
	return retriever.get("www.immoc.com")

运行时才知道传入的retriever有没有get
需要注释来说明接口

c++中的duck typing

template <class R>
string download(const R& retriever) {
	return retriever.get("www.immoc.com");
}
编译时才知道传入的retriever有没有get
需要注释来说明接口

java中的类似代码
<R extends Retriever>
String download(R r) {
	return r.get("www.immoc.com");
}
传入的参数必须实现Retriever接口
不是duck typing

go语言中的duck typing
同时需要Readable, Appendable怎么办? (apache polygene)
同时具有python, c++ 的duck typing的灵活性
又具有java的类型检查
*/
