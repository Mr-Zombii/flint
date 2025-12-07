@.str.a = constant [14 x i8] c"Hello World!\0A\00", align 1

declare external ccc void @printString(i8* %0)

define void @main() {
entry:
        call void @printString(i8* getelementptr ([14 x i8], [14 x i8]* @.str.a, i32 0, i32 0))
        ret void
}