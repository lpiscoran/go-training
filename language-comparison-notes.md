# Using familiar languages for Complex Concepts  


## Struct and map literals

Struct literal: Java is new Vertex(1, 2), and Python is a dataclass call like Vertex(1, 2). Go’s positional form Vertex{1, 2} is Python’s positional args, and the named form Vertex{X: 1, Y: 2} is Python’s keyword args, Vertex(x=1, y=2).
Map literal: Python’s {"alice": 30} is almost identical to Go’s map[string]int{"alice": 30}. In Java, it’s Map.of("alice", 30) or a HashMap you fill with put.
Struct vs. class: A Go struct is like a Java class that has only fields, or a Python dataclass. Methods get attached separately, which is the next lesson.

Why pointers felt confusing: Go copies where you’re used to sharing

In Java and Python, w = v makes both names point at the same object. In Go, w := v copies a struct. That’s why you needed & to get the sharing behavior you’d get for free elsewhere. When Java passes an object to a method, it passes a reference. In Go, you choose between passing a copy (Vertex) and passing a reference (*Vertex).

Maps: where Go differs from both

	Python	Java	Go
Missing key	KeyError	get returns null	returns the zero value
Check for key	"k" in d or d.get	containsKey	v, ok := m["k"]
Iteration order	insertion order	unspecified (HashMap)	deliberately random
Pass to function	shared	shared	shared (maps are reference-like)

Two consequences matter for the WordCount exercise:

Because a missing key returns the zero value, counts[word]++ just works on a new word, since it starts at 0. You don’t need defaultdict or Counter.
Don’t depend on loop order over a map. If you need sorted output, collect the keys into a slice and sort that.

Also, reading from a nil map is fine, but writing to one panics. Always create a map with make(map[string]int) or a literal before writing to it.

A few more mappings for what’s coming

Interfaces: Java makes you write implements. Go doesn’t, because a type satisfies an interface just by having the right methods. That’s closer to Python’s duck typing, but checked at compile time.
Errors: there are no exceptions. Functions return an error value, and you check it with if err != nil, which replaces try/catch.
Generics: the type parameters section will feel familiar if you’ve used Java’s <T>.

When you get to WordCount, the function takes a string and returns a map[string]int. strings.Fields(s) splits on whitespace, which is the same job as Python’s s.split(). Try it, and tell me where you get stuck.

## Methods  
Most of what you know about class methods carries over, with a few differences worth knowing before you start.

The same method in three languages

java
// Java
class Vertex {
    double x, y;
    double abs() { return Math.sqrt(x*x + y*y); }
}
python
# Python
class Vertex:
    def __init__(self, x, y):
        self.x, self.y = x, y
    def abs(self):
        return math.sqrt(self.x**2 + self.y**2)
go
// Go
type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

How the pieces map

Concept	Java / Python	Go
Data	fields inside the class	fields in a struct
Method location	inside the class body	outside, as a separate function
this / self	implicit in Java, first parameter in Python	an explicit receiver, written between func and the method name
Call	v.abs()	v.Abs()

The receiver (v Vertex) is Go’s version of self, except you name it yourself. Convention is a short name, usually the first letter of the type, rather than self or this.

Differences to know up front

A method is just a function with a special first argument. These two do the same thing:
go
   func Abs(v Vertex) float64     // regular function
   func (v Vertex) Abs() float64  // method

The Tour says this explicitly, and it ties back to your scale function. scale(v *Vertex, f int) becomes func (v *Vertex) Scale(f float64), called as v.Scale(10).

Methods can be attached to non-struct types too, like type MyFloat float64. There’s no Java or Python equivalent, since a method needs a class there. The one restriction is that the type must be declared in the same package.
Pointer vs. value receiver is the big one. In Java and Python, this/self is always a reference, so a method can always modify the object. In Go you choose:
go
   func (v Vertex) Abs() float64      // value receiver: gets a COPY
   func (v *Vertex) Scale(f float64)  // pointer receiver: can modify the original

This is exactly your scaleByValue vs. scaleByPointer experiment. A value receiver that tries to modify v.X changes only its own copy.

You can still write v.Scale(10) on a plain Vertex variable. Go quietly turns it into (&v).Scale(10) for you.

No constructors, inheritance, or private keyword.
Constructors: by convention, a plain function like NewVertex(...), as in the NewServer example earlier.
Inheritance: doesn’t exist. Go reuses behavior through composition (embedding one struct in another) and through interfaces.
Visibility: the capitalization rule you already know (Abs exported, abs not). It applies per package, not per type, so there’s no private field that other code in the same package can’t touch.
Interfaces are satisfied implicitly. Java needs implements Shape. In Go, if your type has the methods an interface lists, it satisfies the interface automatically, with no declaration linking them. That’s the second half of this lesson, and it’s closer to Python’s duck typing, but checked at compile time.

Rule of thumb for receivers: if any method on a type needs to modify it, or the struct is large, or it holds something like a mutex, use pointer receivers. And keep a type consistent: don’t mix value and pointer receivers on the same type. In your REST API, nearly every method will be on a pointer receiver, like func (a *API) GetUser(...).

Try this to lock it in. Write both Abs (value receiver) and Scale (pointer receiver) on Vertex in methods/vertex/main.go. Then change Scale to a value receiver and watch the vertex stop changing, which is the same result as scaleByValue earlier.

When you’ve read the first few pages, tell me whether the receiver syntax still feels backwards, because func (v Vertex) Abs() puts the type before the name, and that trips up most people from Java.