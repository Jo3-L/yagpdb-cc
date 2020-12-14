{{/* Reaction CC to configure Bump Pings 

  Trigger:- Added Reactions Only
  
  */}}

{{/*Configuration Values*/}}
{{$roleid := 778219294986207232}} {{/* Update the ROLEID */}}
{{$onetimeroleid := 7896785432678910}} {{/*RoleID of your one time ping role.*/}}
{{/*Configuration Values End*/}}

{{/*If you don't want one time ping role, use this code.
Remember to remove the slash(/) and asterisk(*) in front of every line.*/}}
{{/*CODE STARTS*/}}
{{/*if and (eq .Reaction.Emoji.Name "🔔") (hasRoleID $roleid)*/}}
{{/*removeRoleID $roleid*/}} {{/*sendDM "Bump Pings has been removed from you" */}}
{{/*else if eq .Reaction.Emoji.Name "🔔" */}}
{{/*addRoleID $roleid*/}} {{/*sendDM "Bump Pings has been added to you"*/}}
{{/*end*/}}
{{/*deleteMessageReaction nil .Reaction.MessageID .Reaction.UserID .Reaction.Emoji.Name*/}}
{{/*CODE ENDS*/}}

{{if eq .Reaction.Emoji.Name "🔔"}}
{{addMessageReactions nil .Message.ID "❌" "1️⃣" "♾"}}
{{if eq .Reaction.Emoji.Name "❌"}}
{{deleteAllMessageReactions nil .Message.ID}}
{{addMessageReactions nil .Message.ID "❌"}}
{{else if eq .Reaction.Emoji.Name "1️⃣"}}
{{if hasRoleID $onetimeroleid}}
{{removeRoleID $onetimeroleid}} {{sendDM "One time Bump Ping has been removed from you."}}
{{else}}
{{addRoleID $onetimeroleid}}{{sendDM "One time Bump Ping has been added to you."}}
{{removeRoleID $onetimeroleid 7500}} {{/*Actually 7200, but due to some reasons, 7500. You can change it back to 7200 if needed*/}}
{{end}}
{{else if eq .Reaction.Emoji.Name "♾"}}
{{if hasRoleID $roleid}}
{{removeRoleID $roleid}} {{sendDM "Infinite Bump Pings has been removed from you"}}
{{else}}
{{addRoleID $roleid}} {{sendDM "Infinite Bump Pings has been added to you"}}
{{end}}
{{end}}
{{end}}
