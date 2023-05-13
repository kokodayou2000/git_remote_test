package t1

func Do(name string) string {
	return doSomeThing(name)
}

func doSomeThing(name string) string {
	p := Person{
		Name: "Tom",
		age:  18,
	}

	println(p.Name, p.age)
	return name + "do do do"
}

type Person struct {
	Name string
	age  int
}
