# Custom Workflow Task

- https://cuetorials.com/go-api/workflows/custom/

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
TF:  {
	a: {
		foo:   1
		hello: "world"
		status: {
			bar: 2
		}
	}
	b: {
		foo: 2
		status: {
			bar: 3
		}
	}
	c: {
		foo: 6
		goo: 9
	}
}
TF:  {
	foo: 6
	goo: 9
}
CustomTask: 6 {
	foo: 6
	goo: 9
}
TF:  {
	a: {
		foo:   1
		hello: "world"
		status: {
			bar: 2
		}
	}
	b: {
		foo: 2
		status: {
			bar: 3
		}
	}
	c: {
		foo: 6
		status: {
			bar: 7
		}
		goo: 9
	}
}
```