1 : reserved word: package
1 ID, name=main
3 : reserved word: import
3 (
4 ERROR:"
4 ID, name=os
4 ERROR:"
5 ERROR:"
5 ID, name=flag
5 ERROR:"
6 ERROR:"
6 ID, name=lexgo
6 /
6 ID, name=scanner
6 ERROR:"
7 )
9 : reserved word: func
9 ID, name=main
9 (
9 )
9 {
10 ID, name=fptr
10 :=
10 ID, name=flag
10 ERROR:.
10 ID, name=String
10 (
10 ERROR:"
10 ID, name=fpath
10 ERROR:"
10 ERROR:,
10 ERROR:"
10 ID, name=source
10 ERROR:.
10 : reserved word: go
10 ERROR:"
10 ERROR:,
10 ERROR:"
10 ID, name=file
10 ID, name=path
10 ID, name=to
10 ID, name=read
10 ID, name=from
10 ERROR:"
10 )
11 ID, name=dirptr
11 :=
11 ID, name=flag
11 ERROR:.
11 ID, name=String
11 (
11 ERROR:"
11 ID, name=outputdir
11 ERROR:"
11 ERROR:,
11 ERROR:"
11 ERROR:"
11 ERROR:,
11 ERROR:"
11 ID, name=file
11 ID, name=path
11 ID, name=to
11 ID, name=write
11 ID, name=to
11 ERROR:"
11 )
12 ID, name=flag
12 ERROR:.
12 ID, name=Parse
12 (
12 )
14 : reserved word: if
14 ID, name=fptr
14 ==
14 ID, name=nil
14 {
15 ID, name=os
15 ERROR:.
15 ID, name=Exit
15 (
15 NUM, val=1
15 )
16 }
17 : reserved word: if
17 ID, name=dirptr
17 ==
17 ID, name=nil
17 {
18 ID, name=dir
18 ERROR:,
18 ID, name=err
18 :=
18 ID, name=os
18 ERROR:.
18 ID, name=Getwd
18 (
18 )
19 : reserved word: if
19 ID, name=err
19 ERROR:!
19 =
19 ID, name=nil
19 {
20 ID, name=os
20 ERROR:.
20 ID, name=Exit
20 (
20 NUM, val=1
20 )
21 }
22 ID, name=dirptr
22 =
22 ERROR:&
22 ID, name=dir
23 }
25 ID, name=scanner
25 ERROR:.
25 ID, name=SourcefileWalk
25 (
25 *
25 ID, name=fptr
25 ERROR:,
25 *
25 ID, name=dirptr
25 )
26 }
