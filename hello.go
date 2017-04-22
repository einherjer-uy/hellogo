package main

import (
	"strconv"
	"fmt"
	"time"
	"math/rand"
)

func main() {
	//helloWorld()
	//multipleAssignment()
	//println(createPerson().Age)
	//createAndModify()

	//testComposition()

	//intPointerTest()

	//testArrays()

	//testInterface()

	produce()
}

func randomInt() int {
	//return rand.Int()
	return rand.Intn(250)
}

func helloWorld() {
	var name string
	name = "world"
	fmt.Printf("hello, %s (random: %d)\n", name, randomInt())
}

func multipleAssignment() {
	power := 1000
	fmt.Printf("default power is %d\n", power)
	name, power := "Goku", 9000
	fmt.Printf("%s's power is over %d\n", name, power)
}

func intPointerTest() {
	i := 0
	intPointerModify(&i)
	println(i)
}
func intPointerModify(i *int) {
	*i++ //usar * adelante para obtener el valor apuntado por el puntero
}

type SuperPerson struct {
	Type string
}
func (s *SuperPerson) GetType() string {
	return s.Type
}

type Person struct {
	*SuperPerson //podria ser SuperPerson, en vez de referencia
	Name string
	Age int
	Father *Person
	Logger //composicion con interfaz
}
func (Person) FuncQueNoAccedeAFieldsNoPrecisaVariable() string {
	return "hello soy Person"
}
func (p *Person) SetAge(age int) {
	p.Age = age
}
func (p *Person) GetType() string { //overwrite SuperPerson.GetType
	superType := p.SuperPerson.GetType() // access "component" using the implicit field named the same as the compontent Type
	return superType + "SOMETHING"
}

func NewPerson(name string, age int) *Person {
	return &Person{
		SuperPerson: &SuperPerson {"tipo"},
		Name: name,
		Age: age,
		Father: &Person{
			Name: "father of " + name,
			Age: 30 + age,
			Father: nil,
		},
		Logger: LoggerImpl{}, //proveo la implementacion de la interfaz
	}
}
func createPerson() Person {
	person := Person{Name: "Pepe", Age: 10}
	return person
}

func modifyPerson(p Person) {
	p.Age += 1
	p.SetAge(p.Age + 10) //puedo llamar a SetAge a pesar de q el "receiver" del metodo SetAge sea un *Person y no Person
}

func modifyPersonPointer(p *Person) {
	p.Age += 1
	p.SetAge(p.Age + 10)
}

func createAndModify() {
	person := createPerson()
	modifyPerson(person)
	println(person.Age)

	var p *Person // p es de tipo "puntero a"
	p = &person // a p le asigno una "referencia a" person
	println(p)

	modifyPersonPointer(p)
	println(p.Age)

	p2 := NewPerson("Pep", 8)
	println(p2)
}

func testComposition() {
	//person := createPerson()
	//print(person.GetType()) //null porinter porq createPerson no inicializa completamente la person

	person2 := NewPerson("name", 10)
	print(person2.GetType())
}

//COLLECTIONS

func testArrays() {
	var a [10]int
	a[0] = 100
	print(len(a))

	b := [3]int {1, 2, 3}
	for index, value := range b {
		println("value " + strconv.Itoa(value) + " at position " + strconv.Itoa(index))
	}
}

//INTERFACES

type SuperLogger interface {
	SuperLog(message string)
}
type Logger interface {
	SuperLogger //composicion de interfaces
	Log(message string, code int)
}

type LoggerImpl struct {

}
func (LoggerImpl) SuperLog(message string) {
	println(message)
}
func (LoggerImpl) Log(message string, code int) {
	println(strconv.Itoa(code) + "-" + message)
}

func testInterface() { // ver tb arriba el struct Person que se compone de un Logger
	var l Logger
	l = LoggerImpl{}
	l.SuperLog("test message")
	l.Log("test message", 1)
}

// CONCURRENCY

type Worker struct {
	id int
}
func (w *Worker) process(c chan int) {
	for {
		data := <-c
		fmt.Printf("worker %d got %d\n", w.id, data)
		time.Sleep(time.Millisecond * 3000) // simulate expensive processing
	}
}

func produce() {
	c := make(chan int, 100) //buferred channel, len 100
	for i := 0; i < 5; i++ {
		worker := &Worker{id: i}
		go worker.process(c)
	}
	for {
		time.Sleep(time.Millisecond * 50)

		select { // resumen
				// si no tiene ni timeout ni default bloquea forever (una vez q se lleno el buffer)
				// si tiene timeout espera el timeout antes de responder (una vez q se lleno el buffer)
				// si tiene default se va por el default (una vez q se lleno el buffer)
			case c <- rand.Int():
				fmt.Println("Channel length is " + strconv.Itoa(len(c)))
			case <-time.After(time.Millisecond * 1000):
				fmt.Println("timed out")
			default:
				println("dropped")
		}
	}
}

