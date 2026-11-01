# Using familiar languages for Complex Concepts  


Struct and map literals

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