// https://playground.ponylang.io/?gist=3fad48120e1e1f2263970c96c4c36de2
// https://tutorial.ponylang.io/types/classes.html

class FooIso
  var s: String ref
  
  new iso create(s': String iso) =>
    s = consume s'

  fun string(): String iso^ =>
    s.clone()

actor Main
  new create(env: Env) =>
    let e = FooIso("hi".clone())
    let s = recover iso String end
    // s can be modified before placing it in e
    s.append("stuff")
    e.s = s.clone()
    
env.out.print(e.string())
