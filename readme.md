# Exercise

The public repository for this project can be found [here](https://github.com/b9lab/ida-exercise-week-2-student-repo). You may want to clone it to your local computer for practice.

## Preparation

We created a Docker container with the included `Dockerfile`:

```sh
$ docker build . -t exercise-w2
```

## What we set up

We created this repository with Ignite CLI v0.22.1 and the following command:

```sh
$ docker run --rm -it -v $(pwd):/exercise -w /exercise exercise-w2 ignite scaffold chain github.com/b9lab/other-world
```

The idea behind it is that this is the blockchain backing a future metaverse. The blockchain will account for anything of value in it.
