# z

`z` exposes a curated set of utility functions.

## Goals 

Its primary goals are two-fold:

1. Make the best commonly used utilities discoverable; "Ugh! What's the import for that helper again?"
2. Reduce hand strain; "I can finally stop writing `func must(...)` everywhere!"

## Why name it 'z'?

Because.

Besides that, the choice of name is driven by:

- Brevity: a smaller package name is more ergonomic (`z.P("foo")` vs. `utils.P("foo")`)
- Collision resistance: 'z' is _probably definitely_ less likely to be shadowed by names in dependent packages 
- Recognizable: 'z' is memorable (opinion)

