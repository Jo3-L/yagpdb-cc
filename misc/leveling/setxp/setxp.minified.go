{{$n:=reFind `(?i)xp|level` .Cmd}}{{$k:=0}}{{$d:=0}}{{with .CmdArgs}}{{$k =index . 0 | userArg}}{{end}}{{if and (eq $n "level") (eq (len .CmdArgs) 2)}}{{$g:=$.nil}}{{$h:=false}}{{with reFind `\A\d+\z` (index .CmdArgs 1)}}{{$h =true}}{{$g =toInt .}}{{end}}{{if and $h $k}}{{$b:=mult 100 (mult $g $g)}}{{$p:=dbSet $k.ID "xp" $b}}{{$d =$g}}{{printf "Successfully set **%s**'s level to %d!" $k.String $g}}{{else}}
The syntax for this command is `-setlevel <user> <level>`. 
{{end}}{{else if eq (len .CmdArgs) 2}}{{$o:=$.nil}}{{$h:=false}}{{with reFind `\A\d+\z` (index .CmdArgs 1)}}{{$h =true}}{{$o =toInt .}}{{end}}{{if and $h $k}}{{$i:=0}}{{with (dbGet $k.ID "xp")}}{{$i =.Value}}{{end}}{{$e:=roundFloor (mult 0.1 (sqrt $i))}}{{$p:=dbSet $k.ID "xp" $o}}{{$c:=roundFloor (mult 0.1 (sqrt $o))}}{{if ne $e $c}}{{$d =$c}}{{end}}{{printf "Successfully set **%s**'s XP to %d!" $k.String $o}}{{else}}
The syntax for this command is `-setxp <user> <xp>`.
{{end}}{{else}}
The syntax for this command is `-setxp/setlevel <user> <level|xp>`.
{{end}}{{if $d}}{{$a:=sdict "type" "stack"}}{{with dbGet 0 "roleRewards"}}{{$a =sdict .Value}}{{end}}{{$l:=$a.type}}{{$j:=0}}{{$m:=-1}}{{range $g, $f:=$a}}{{if and (le (toInt $g) (toInt $d)) (or (gt $m (sub (toInt $d) (toInt $g))) (eq $m -1)) (eq $l "highest")}}{{$m =sub (toInt $d) (toInt $g)}}{{$j =$f}}{{end}}{{if and (ge (toInt $d) (toInt $g)) (not (targetHasRoleID $k $f)) (eq $l "stack") (ne $g "type")}}{{giveRoleID $k $f}}{{else if and (targetHasRoleID $k $f) (eq $l "highest")}}{{takeRoleID $k $f}}{{end}}{{end}}{{if $j}}{{giveRoleID $k $j}}{{end}}{{end}}