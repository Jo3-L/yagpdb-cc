---
sidebar_position: 11
title: parseArgs
---

This snippet provides a reusable template which you can add to your custom commands.  
It separates predefined flags from positional arguments within input; for example:
- `-command -m 123 positional arg -f`
    - might be parsed into:
    
    ```
    positional: ["positional" "arg"]  
	  flags: map {  
	    "m" => "123"  
	    "f" => true  
	  }
	  ```

A more elaborate example is shown in the code below.

```go
{{/* The 'define' block below is all you need to copy into your command; everything else is part of the example code.' */}}
{{define "parseFlags"}}
	{{.Set "Out" (sdict
		"Positional" (cslice)
		"Flags" (dict)
	)}}

	{{$curFlag := ""}}
	{{$lastIdx := sub (len .Args) 1}}
	{{range $i, $arg := .Args}}
		{{- if $curFlag}}
			{{- $.Out.Flags.Set $curFlag $arg}}
			{{- $curFlag = ""}}
		{{- else if and ($id := $.Flags.Get $arg) (ne $i $lastIdx)}}
			{{- $curFlag = $id}}
		{{- else if $id := $.Switches.Get $arg}}
			{{- $.Out.Flags.Set $id true}}
		{{- else}}
			{{- $.Out.Set "Positional" ($.Out.Positional.Append $arg)}}
		{{- end -}}
	{{end}}
{{end}}

{{/* Example usage, to parse a command similar in structure to the 'rolemenu create' built-in command */}}

{{$query := dict
	"Flags" (dict
		"-m" "MessageID"
		"-msg" "MessageID"

		"-skip" "Skip"
		"-s" "Skip"
	)

	"Switches" (dict
		"-nodm" "NoDM"
		"-rr" "RemoveRole"
	)

	"Args" .CmdArgs
}}
{{template "parseFlags" $query}}

{{/*
	Access flags by their ID (the value in the map supplied to parseFlags).
	The value will be nil if the user did not supply that flag; otherwise it will
	be the argument directly after the flag. For example, $msgID would have the value "1"
	if the input was "-m 1".
*/}}
{{$msgID := $query.Out.Flags.MessageID}}

{{/*
	Same thing for switches. In this case, the value will be nil if the user did not supply
	that switch; otherwise it will be 'true'.
*/}}
{{$nodm := $query.Out.Switches.NoDM}}

{{/*
	Finally, we can access any excess arguments that are neither flags nor switches by indexing
	into the Positional slice. In this case, $first would have the value "a" if the input was
	"a -nodm -m 123".
*/}}
{{$first := index $query.Out.Positional 0}}
```
