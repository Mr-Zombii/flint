@.str.a = constant [6 x i8] c"Hello\00", align 1
@.str.b = constant [5 x i8] c"true\00", align 1
@.str.c = constant [6 x i8] c"false\00", align 1

declare external ccc void @print(i8* %0)

define i8* @test(i64 %i) {
entry:
        %0 = alloca i64
        store i64 %i, i64* %0
        %1 = load i64, i64* %0
        %2 = icmp eq i64 %1, 1
        br i1 %2, label %match.0.arm.0, label %match.0.next.0

match.0.arm.0:
        br label %match.0.merge

match.0.next.0:
        br i1 true, label %match.0.arm.1, label %match.0.next.1

match.0.arm.1:
        %3 = load i64, i64* %0
        %4 = icmp eq i64 %3, 10
        %5 = icmp eq i1 %4, true
        br i1 %5, label %match.1.arm.0, label %match.1.next.0

match.0.next.1:
        br label %match.0.merge

match.1.arm.0:
        br label %match.1.merge

match.1.next.0:
        %6 = icmp eq i1 %4, false
        br i1 %6, label %match.1.arm.1, label %match.1.next.1

match.1.arm.1:
        br label %match.1.merge

match.1.next.1:
        br label %match.1.merge

match.1.merge:
        %7 = phi i8* [ getelementptr ([5 x i8], [5 x i8]* @.str.b, i32 0, i32 0), %match.1.arm.0 ], [ getelementptr ([6 x i8], [6 x i8]* @.str.c, i32 0, i32 0), %match.1.arm.1 ], [ undef, %match.1.next.1 ]
        ret i8* %7

match.0.merge:
        %8 = phi i8* [ getelementptr ([6 x i8], [6 x i8]* @.str.a, i32 0, i32 0), %match.0.arm.0 ], [ %7, %match.0.arm.1 ], [ undef, %match.0.next.1 ]
        ret i8* %8
}

define void @main() {
entry:
  %0 = call i8* @test(i64 3)
  call void @print(i8* %0)
  ret void
}


