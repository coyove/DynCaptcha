# DynCaptcha
## A simple dynamic captcha written in Go

### Usage

```
buf, answer := DynCaptcha.New(0)
```

`buf` is the generated gif image in `[]byte`. `answer` is an `int` holding the answer of this captcha. `0` is the random seed.

### Demo

Find the red circle's final position.

<img src="https://github.com/coyove/DynCaptcha/blob/master/demo.gif" height="256" width="256" >

It should be 21.
