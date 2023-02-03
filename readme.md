# FN

## summary

function utilities for go 1.18+

## abilities

Panic: check if returns error then panic,otherwise return values.

Panics##: same a Panics but with closure ability and returns a Runnable drop all results.

Recovers: wrap functions with recovery.

SliceXXX: slice operates.

List: object style slice.

MapXXX: map operates.

Map : object style map.

Drop###: drop the first N result.

DropLast###: drop the last N result.

Clos###: closure with first N argument.

ClosLast##: closure with last N argument.

## note

1. all generator are written in fn_gen_test.
2. not full test coverage, use as your own risks.
3. not real stable API.

## license

MIT