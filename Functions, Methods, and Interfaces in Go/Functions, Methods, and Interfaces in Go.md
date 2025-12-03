why use functions
a set of instructions with a name

func main(){
    fmt.Printf("Hello,world.")

}

 function declaration

 why use?
 reysability

 good for commonly used operations

 abstruction
details are hidden in the function

all you need to do is just call the function

only need to understand operations

naming is important

#function parameters and return values
parameters are listed in parenthesis after function name

arguments are supplied in the call
func foo(x int,y int){
    fmt print(x*y)
}

func main(){
    foo(2,3)
}
---
func (x,y int)

return values
type of return value after paraeters in declaration

foo(x int) int {
    return x+1
}

两个返回? python是否有类似语法？
func foo(x int)(int,int){
    return x,x+1
}
a,b := foo(3)

# call by value,reference

passed arguments are copied to parameters

modifying parameter will not influe variable

advatage: data encapsulation

programmer can pass a pointer as an argument
---
foo(y *int){
    y* = *y+1
}

func main(){
    x:=2
    foo(&x)
    fmt print(x) 3
}
---
dont need to copy arguments，
比如变量很大，几万个元素，就不需要拷贝，直接给地址；

disadvantage here: data encapsulation

- passing arrays and slices

array arguements are copied
array can be big so this can be a problem

func foo(x [3]int) int{
    return x[0]
}

main(){
    a:=[3]int{1,2,3}
    print(a)
}

----
fuc foo(x *[3]int) int{
    (*x)[0] = (*x)[0]+1 // 这里为什么要用括号包裹x？
}
func main(){
    a:=[3]iint{123}
    foo(&a)
    fmt.Print(a)
}

pass slices instead
slice container a potier to the array
passing a slice copies the pointer
func foo(sli int)int{
    sli[0] = sli[0] + 1
}
func main(){
    a := []int{1,2,3}
    foo(a)
    print(a)
}

try to use slice instead of array in go 
! ★

well=written functions
code is functions and data
if you are asked to find a freatuure, you can find it quickly

understandability
debugging principles

1.function is written incorrectly
2. data that the function uses is incorrect.

supporting debugging

data needs to be traceable
determine if actual behavior matches desired behavir

- Guidelines for functions

function naming
give functions a good name
parameter naming counts too


merging behaviors make code  complicated :bad

reducing parameter Number

functions shoud be simple 
function call hierarchy

control flow complexity

