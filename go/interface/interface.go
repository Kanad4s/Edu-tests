package main

import "fmt"

type BigAPIClient struct {
	id   int
	name string
}

func (c *BigAPIClient) Print() {
	fmt.Println(c.name)
}

func (c *BigAPIClient) FetchMessages() (msg string) {
	msg = fmt.Sprintf("id %d, name: %s", c.id, c.name)
	return
}

func (c *BigAPIClient) SetId(id int) {
	c.id = id
}

type TestInterface interface {
	Print()
	FetchMessages() (msg string)
	SetId(id int)
}

type Company struct {
	array []*TestInterface
}

func (c *Company) AddIntrfc(testInterface TestInterface) {
	c.array = append(c.array, &testInterface)
}

func (c *Company) TestCompany() {
	for _, val := range c.array {
		testFunc2(*val)
	}
}

func testFunc2(testInterface TestInterface) {
	testInterface.Print()
	fmt.Println(testInterface.FetchMessages())
	testInterface.SetId(20)
	fmt.Println(testInterface.FetchMessages())
	testInterface.SetId(200)
	testInterface.Print()
}

func testFunc(testInterface TestInterface) {
	testInterface.Print()
	fmt.Println(testInterface.FetchMessages())
	testInterface.SetId(10)
	fmt.Println(testInterface.FetchMessages())
	testInterface.SetId(100)
	testInterface.Print()
}

type Person struct {
	id   int
	name string
}

func (c *Person) Print() {
	fmt.Println(c.name)
}

func (c *Person) FetchMessages() (msg string) {
	msg = fmt.Sprintf("Person:\nid %d, person: %s", c.id, c.name)
	return
}

func (c *Person) SetId(id2 int) {
	c.id *= id2
}

func main() {
	client := &BigAPIClient{id: 5, name: "apiClient"}
	testFunc(client)
	fmt.Println(client.id)

	company := Company{}
	company.AddIntrfc(client)

	person := &Person{id: 9, name: "person"}
	testFunc(person)

	company.AddIntrfc(person)

	company.TestCompany()
	fmt.Println(client.id)
}
