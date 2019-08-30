# Tortoise

A port of ShellJS to Go, providing familiar and powerful commands in a platform agnostic library.

```
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNKOxxkxxkO0KNWMWWWWWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWNKOkOOOOOkO0Odc;:cc;,,;;;:coddddxkKWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMN0dl:coddol:;;;:clllllllcllc:,';cc::;:ox0NMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWWNKko:;:codddolc:;;;:cllcclooooddl;;:;;;;:cccdk0NMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOxkxdoccccloddddddxdl;:clddol:;;:ldxxdlcc;,..',:ccldkKXNMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOl::oocloooccllolccoxkklcldxxddo:,';ldxxxdl;...,..,;colccldONMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWXOl::cloooddlcldkkxdc:ldkkxoodddoddc,'';ldxxdl;'..,,..':lol:,.'l0WMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKkocccclolooolloxkkxxo:;:oxkkkkddlcloc,..':odddo:;'..,,.';cool:,,;cxXMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKxlc:cclllcccclllodddoll:,;ldxkOOxoc;:c::,..,loodoc,'..',,..;looc;',::l0WMMMMMMMMMMMM
MMMMMMMMMMMWWWWWWMMMMMMMMMMMMMMWNK0oclooolc:;:c::cloddxdoll::;coxxkkxo;';;;;,...;lllll;'........;lllc,',:::dXWMMMMMMMMMM
MMWX0OkkkkxxxxxxkkO0KXNWMMMWX0OkddoodkOOOklc:::::::coodolcc::;,cdxxxdl;',,'''...':llll:'.........;cll:,',c:;ckNMMMMMMMMM
MKdc;,,,,',:ooc:cc::ldddk0OxodolodxkOOOkxxddlc:;:cc::clc::::;,,,cdddo:,..........'clllc,..........,cll:,';;,,;xNMMMMMMMM
MO;''''...';c:,''',;clcclll:lxxxkOOOOkxoodxO0Okdoll::cc;;;;,''..;oddd:'...........'lool:'.....'''';looc::llc:clxKWMMMMMM
M0:''',;;;;,,;,'..';lc,,,'''cxkkkkkxxxdxOOO0000Oxxkxddlc:ccc::::cdkxxoc::;,,,'',;;coodl:cllclloc;;;:loc;,:cllccloONMMMMM
MNOc'''',;;;;;,...,:c;,'...';cokOxdxxxk0KK0000kdloxxkkxocoxxxddxxxddddooooooolcloollcc:'';::ccl:,,,,,:c:'';:lolcldkKWMMM
MMMXkdllcc:::;;;;;;::,''',,;:::x0kxkOO0KKKK0Oxdolodxxddl::lollloooccccc:;:c:c:,;;;;;;:,..',;::c:...'',:c,..';clllddxOKXN
MMMMMWWNX0kdoolc;,'',;,'''',,,,;d00K0OOOO0Okdolcclodxxdc;;:ccloollol;,,,,,;::,..',,;;;,...,;;;::....'''::'...'',;ldxdxON
MMMMMMMMMMMWNNXOol:;''''',;;;,'.':xKKkxkxkkdolc::ccldxdl:;;;;:cc:lol,'.',,,;:;'...'',;;....'',;:'.......';;,'....';lkXWM
MMMMMMMMMMMMMMWOolc:;:cloooooooc'..'::cldkkxool:::ccodo:;,,;,;;;;clc,...''',::,.....',;'......,;;'........;:::,'..,xNMMM
MMMMMMMMMMMMMMNkl:,,,;::;''',:oxc. .';:clloddoccclc:odl:,'.''''',:c;'......';;'.......,'.......';:;... ...';::;:cd0WMMMM
MMMMMMMMMMMMMMMXxl:;,:c;,,,:ldOKKdccodddd:'''',;::;;ccc;,'......';:;'.......,;''.......,'.......,;;,'','......,o0WMMMMMM
MMMMMMMMMMMMMMMMX0xolc:;,,,clcoOXKdllooddol:;'.....',,'.........';;,,''....';:,,'......,,'.'',,,:::;;'...      ,0MMMMMMM
MMMMMMMMMMMMMMMMMWNKOo;'.,cdxdolodlcccccllllllc;'......''''..';::;,'',;,,;:c:;,,,','';cccccllc;.....      .'.. .kMMMMMMM
MMMMMMMMMMMMMMMMMMMMWNOdkKWMMWN0kxdc:c:ccllolodlcdxoc;;;;::cccccc;;;,,,;;:cc::;;;;;;;::;;,'...':lc'  .....:;...;0MMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMW0o:cc::cllcccllcc:::llc:::;,,'''''''''''''''''''.'''';clxOKNWMWo  ',..'clcokXWMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNOxc,;;,:ccc:';lkkdolc:,';cooddoooooddddxxxxkkkkOO0KNWMMMMMMMMKl,;d0KKKKNMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWKxl:;;;,;:;,;cdxoc;;,,:kNMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMWX0XMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMX0Odcc:,;:::codooolccd0NMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMNXXKKXKKXXNNNNNNXXNWMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
MMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMMM
```

## Why?

Partly because a need arose, but also for fun! I started programming in JavaScript before I ever even knew what Go was and I'm a big of shell scripting to help automate tasks. As I became more and more familiar with Go and abandoned Node entirely, I found it difficult to phase out usage of gulp / ShellJS in part because no Go library (to my knowledge) offered simple access to utilities such as `sed` that work on any platform. I also thought this would be a great way to familiarize embolden my Go knowledge, and hopefully get some feedback on my Go code quality.

## Why is it called tortoise?

Because `gopher-tortoise-shell` is too long
