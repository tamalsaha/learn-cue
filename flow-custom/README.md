
```bash
> go run flow-custom/main.go
Custom Flow Task
TF:  {
	a: {
		foo:   1
		hello: string
		status: {
			bar: int
		}
	}
	b: {
		foo: 2
		status: {
			bar: int
		}
	}
	c: {
		foo: a.status.bar * 3
		goo: b.status.bar * 3
	}
}
TF:  {
	foo:   1
	hello: string
	status: {
		bar: int
	}
}
TF:  {
	foo: 2
	status: {
		bar: int
	}
}
TF:  {
	foo: a.status.bar * 3
	goo: b.status.bar * 3
}
TF:  a.status.bar * 3
TF:  b.status.bar * 3
CustomTask: 2 {
	foo: 2
	status: {
		bar: int
	}
}
CustomTask: 1 {
	foo:   1
	hello: string
	status: {
		bar: int
	}
}
TF:  {
	a: {
		foo:   1
		hello: string
		status: {
			bar: int
		}
	}
	b: {
		foo: 2
		status: {
			bar: 3
		}
	}
	c: {
		foo: a.status.bar * 3
		goo: 9
	}
}
TF:  {
	foo: a.status.bar * 3
	goo: 9
}
TF:  a.status.bar * 3
TF:  9
Error: tasks.c.goo: cannot use value 9 (type int) as struct
```