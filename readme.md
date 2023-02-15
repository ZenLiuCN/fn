# FN

## summary

function utilities for go 1.18+

## abilities

`Panic`: check if returns error then panic,otherwise return values.

`Panics##`: same a Panics but with closure ability and returns a Runnable drop all results.

`Recovers`: wrap functions with recovery.

`SliceXXX`: slice operates.

`List`: object style slice.

`ListEx`: object style slice with extra mapping method support.

`MapXXX`: map operates.

`Map`: object style map.

`MapEx`: object style map with extra mapping method support.

`Drop###`: drop the first N result.

`DropLast###`: drop the last N result.

`Clos###`: closure with first N argument.

`ClosLast##`: closure with last N argument.

`Error` & `Error`: create error with file info 
(control by build tag `NormalErr` `FullErr`), 
`Panic` & `Panics` & `Recover` also respect those tags unless manual change exposed `Packer` variable **new in 0.1.10**

## note

1. all generator are written in fn_gen_test.
2. not full test coverage, use as your own risks.
3. not real stable API.

## license

MIT