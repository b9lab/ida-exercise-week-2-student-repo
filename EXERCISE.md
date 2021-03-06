# Student exercise

This repository was created with Ignite CLI v0.21.2 and the following command:

```sh
$ ignite scaffold chain github.com/b9lab/cosmos-academy-exercise-simple
```

The idea behind it is that this is the blockchain backing a future metaverse. The blockchain will account for anything of value in it.

## Hidden files

There are hidden files that:

1. Declare and implement a `WorldParams2` type.
2. Test this type against a future type you have to create.

## Student actions

Your goal is to run a Ignite command to create a new Protobuf type that follows this description:

1. The name is `WorldParams`.
2. It has 3 params exactly, no more, no less, and in this order exactly, with the Protobuf indices going from `1` to `3`:
   1. A `string` for the world's `name`.
   2. A `uint` for the world's `gravity`.
   3. A `uint` for the world's `landSupply`.
3. It is unique in the blockchain store. At what key it is stored does not matter, but you should keep the naming conventions of keeper functions chosen by Ignite.
4. It should not have any associated messages.

If you create any object named `WorldParams2` or try to _fix_ something that you believe is broken with `WorldParams2` that is OK here. When you are in Hackerrank, you should **not** create any object named `WorldParams2` or try to _fix_ something that you believe is broken with `WorldParams2`.

## How we test

We have a hidden test file that expects your output to be properly named and prepared. Based on that:

* If you have a compilation error, it means you have named elements wrongly.
* If the test fails, you have numbered things wrongly in Protobuf, or you have messages for `WorldParams`.
