package main

// Abstraction
//  decoupling behaviour from the implementation details
//  The Unix file system API is a great example of effective abstraction
// 	Roughly five basic functions hide all messy details:
// 		open
// 		close
// 		read
// 		write
// 		ioctl
//  Many different OS things can be treated like files

// Encapsulation
//  hiding implementation details from misuse
//  It's hard to maintain absraction if the details are exposed:
//		The internals may be manipulated in ways contrary to the
// 		concept behind the absrtaction
// 		Users of the abstraction may come to depend on the internal
// 		details - but those might change
// 	Encapsulation usually means controlling the visibility of names
// 	("private" variables)

// Polymorphism
// 	literally means "many shapes" - multiple types behind single interface
// 	Three main types are recognised:
// 		ad-hoc: typically found in function/ operator overloading
// 		parametric: commonly known as "generic programming"
// 		subtype: subclasses substituting for superclasses
// 	"Protocol-oriented" programming uses explicit interface types,
// 		now supported in many popular languages (an ad-hoc method)
// 	In this case, behaviour is completely separate from implementation,
// 		which is good for abstraction

// Inheritance
//  has conflicting meanings:
// 		substitution (subtype) polymorphism
// 		structural sharing of implementation details
//  In theory, inheritance should always imply subtyping:
// 		the subclass should be a "kind of" the superclass

// Inheritance Problems:
//  it injects a dependence on the superclass into the subclass:
// 		what if the superclass changes behaviour?
// 		what if the abstract concept is leaky
// 	Not having inheritance means better encapsulation & isolation
// 	"Interfaces will force you to think in terms of communication
// 	between objects" - Nicolo Pignatelli in Inheritance is evil
// 	See also "Composition over Inheritance" and "Inheritance Tax"

// Alan kay on OOP
// 	"OOP to me means only messaging, local retention and protection
// 	and hidingof state-process, and extreme late binding of all things"
// The aim of the above statement being to:
// 		De-emphasize inheritance hierarchies as a key part of OOP
// 		Emphasize the idea of self-contained objects sending messages
// 		to each other
// 		Emphasize polymorphism in behaviour

// Go offers four main supports for OOP
// 	encapsulation using the package for visibility control
// 	abstraction & polymorphism using interface types
// 	enhanced composition to provide structure sharing

// Go does not offer Inheritance or Substitutability based on types
// 	Substitutability is based only on interfaces: purely a
//  function of abstract behaviour

// Not having classes can be liberating
// 	Go allows defining methods on any user-defined type,
// 		rather than only a "class"
// 	Go allows any object to implement the method(s) of an
// 		interface, not just a "subclass"

func main() {

}
