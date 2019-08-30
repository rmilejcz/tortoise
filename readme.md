# Tortoise

A port of ShellJS to Go, providing familiar and powerful commands in a platform agnostic library.

```
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNXKKKXXNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWXXXXXXXNWNOdlcclc:cllodk0K0O0KNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOxoloddolccllccccllc:;::;,',:clcccdOXWMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKkl;:codddl:;;,,;cllloooooodo:,;::;;:::ldOXWMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNKK0Oxl::cldxxddddoc;;:cllcc::clodxdlcc:;,',;cccox0NWMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMXxlcoolloolllloolloxkxl:coxxdol;',:oxxxdoc;'..'.':cllodkOXWMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKxc:cloolodlcloxxoc:coxkxllodxdddc,'':oxxxdl;...;'..,colc;,;oKWMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNKklcccclooodoloxOOkxoc:lxkkxxddolloc,..'codddl;,..',..':lol:,',:kNMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN0xlcccllllclllclddddolc:;:dxkOOxdl::c:;'..;loddo:,'..,;'.,clol:,,::oKWMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNXxlccccclc::c:cllodddollc:;:oxxkkkdc,;:;;,..':ooooc,'...''..,cllc;',::cxXMMMMMMMMMMM
MMMWXKK0000OOOOO0KXNWWMMMMMWNX0Oxoloxkkxdc:::::::codddolc::;;:oxxxxo:'',,,'...,cllll;'....'...,cllc,',::;lONMMMMMMMMM
MXkoc::;;;;cooccllcoxkO0XX0kxdllodxkOOkkkdlc;;:::::coocc:::;,,cdddoc;..........,clll:'.........,:llc,.,;;,;dNMMMMMMMM
MO;,''....':c:,',,;:llclloccdxdxkOOOOkdooxkkxdllcc:cl:;:;;;,'',ldddc,...........,llll;.........':ool;;:c:;:lxXWMMMMMM
MO;''',,',,,;;'...'cl:;;,,,:xkkOkkkxxddxkO000KOddxdooc:::::;;,,lxkxo:;,,'.....'':oool::::::cc;;:clol:;:lolccld0WMMMMM
MXd;''',;;,,;,'...,cc,,'...,ldkOxddxxk0KK0000kdodkkkxdlldxxddddxxxdddoooollcclooooll:,;:cllll;,,,,:c:,';clolcldONMMMM
MMW0oc::::::;;;,,;::;,''',;:::x0kdxkO0KKK00Okxoldxxxddc:loollodolllll:::ccc:;;::;:::,..,;::cc,..''';c;''';collodx0XNW
MMMMWNXX0kdlccc:;,,;;;'''',;;,:x0000OO0000kxolcloddxxo:;:ccllollll:;,,,,;::;..',,;;;'...,;;:c;....'';c,...';,:ldxxxOX
MMMMMMMMMWNXXKOdc:;''''.',,,''.'lOXKOxkkkkdolc:clloxxdc;;;;;clccoo;'..'',;:;...''',;,....',;:;......',;;,'....':lx0NW
MMMMMMMMMMMMMMNkolc:;:clloolll:'..;llloxkkdooc::cclddl:;,;;;;;;cll;'''''',::,.....',;......';;,........;::;,...'oXMMM
MMMMMMMMMMMMMMKdc:,,,:c:;,,,;ldc. ..,;coodxdoccclccooc;''''''',;c:,......';;'......','......',:;'... ..';::;;;lkXMMMM
MMMMMMMMMMMMMMNOoc;,;c:,,';cox00d::loodo;',,,;:c:;:lc:;,.......,:;,.......,;'.......',.......';;,''''....'.,lkNMMMMMM
MMMMMMMMMMMMMMMWKkolcc;,,;coloOXKxllooddoc;,......,,,..... ....;:;,'......;:,,'......,,....'';:c::;'...     'kWMMMMMM
MMMMMMMMMMMMMMMMMNKOd:,'.,lddlldxoccccllllllc;'.......'''...';::,,,,;,,,:::;,,''''';:cccclcc;'....     .'.. .dWMMMMMM
MMMMMMMMMMMMMMMMMMMWNOddOXWMWX0kxdc:c:ccloolodllool:;;;;:::cccc:;,,,,,;:ccc:;;;;;;;:c:;;,'..';c;. .....,:'..'kWMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOo:cc::cllcccoolc::clc::;:;,,,''''',,,'''''''''''..',:ldk0XWM0' .,...;l:cd0WMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOdc,;;;cccc:,:okxdolc;,,:cloolllllooooddddddxxxkOKXWMMMMMMMNx,,cx000KXWMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKdc:;;;,;:;;:oxdl:;,,,oXWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN00NMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWX0koc:;,;:::coooolcclkXWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNXKKKK0KKXNNNNNXXXWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
```

## Why?

Partly because a need arose, but also for fun! I started programming in JavaScript before I ever even knew what Go was and I'm a big of shell scripting to help automate tasks. As I became more and more familiar with Go and abandoned Node entirely, I found it difficult to phase out usage of gulp / ShellJS in part because no Go library (to my knowledge) offered simple access to utilities such as `sed` that work on any platform. I also thought this would be a great way to familiarize embolden my Go knowledge, and hopefully get some feedback on my Go code quality.

## Why is it called tortoise?

Because `gopher-tortoise-shell` is too long
