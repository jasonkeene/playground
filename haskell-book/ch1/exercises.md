## Exercises: Equivalence

1. λxy.xz

    Not a since z and y need to be distinct
    b since m can be replaced for x and n for y
    Not c since the second argument should not be in the body

    Answer: b

2. λxy.xxy

    Not a since p is not the same as m
    Not b since the body requires three terms
    c since we just replaced x with a and y with b

    Answer: c, but I am a bit confused since the notation is missing a '.'
            (dot) to indicate the body

3. λxyz.zx

    Not a since we are missing the x term in the body
    b since we just replaced z with t and z with s
    Not c since y doesn't exist in the body

    Answer: b

## Chapter Exercises

### Combinators

1. yes
2. no
3. yes
4. yes
5. no

### Normal Form / Divergent

1. normal
2. divergent
3. normal

### Beta Reduction

1. (λabc.cba)zz(λwv.w)

    (λa.λb.λc.cba)zz(λw.λv.w)
    (λc.czz)(λw.λv.w)
    (λw.λv.w)zz
    z

2. (λx.λy.xyy)(λa.a)b

    (λx.λy.xyy)(λa.a)b
    (λy.(λa.a)yy)b
    (λa.a)bb
    bb

3. (λy.y)(λx.xx)(λz.zq)

    (λy.y)(λx.xx)(λz.zq)
    (λx.xx)(λz.zq)
    (λz.zq)(λz.zq)
    (λz.zq)q
    qq

4. (λz.z)(λz.zz)(λz.zy)

    (λz.z)(λz.zz)(λz.zy)
    (λz.zz)(λz.zy)
    (λz.zy)(λz.zy)
    (λz.zy)y
    yy

5. (λx.λy.xyy)(λy.y)y

    (λx.λy.xyy)(λy.y)y
    (λy.(λq.q)yy)y
    (λq.q)yy
    yy

6. (λa.aa)(λb.ba)c

    (λa.aa)(λb.ba)c
    (λb.ba)(λb.ba)c
    (λb.ba)ac
    aac

7. (λxyz.xz(yz))(λx.z)(λx.a)

    (λx.λy.λz.xz(yz))(λx.z)(λx.a)
    (λy.λz.(λx.z)z(yz))(λx.a)
    λz.(λx.z)z((λx.a)z)
    λz.(λx.z)z((λx.a)z)
    λz.z(λx.a)z
    λz.za
