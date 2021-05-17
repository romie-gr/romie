# How to use romie

## Search for a game

Let's search for a game that includes the word "Crash" in its title:

```bash
$ ./romie-mac search -t "Crash"
```

Output:

```bash
INFO[0000] Using configuration file: /Users/drpaneas/.romie/config.yml
    playstation | Crash Bandicoot [SCUS-94900]
    playstation | Crash Bandicoot 3 - Warped [SCUS-94244]
    playstation | Crash Team Racing [SCUS-94426]
    playstation | Crash Bandicoot 2 - Cortex Strikes Back [SCUS-94154]
    playstation | Crash Bash [SCUS-94570]
    playstation | Crash Bandicoot
    playstation | CTR - Crash Team Racing (Europe) (En,Fr,De,Es,It,Nl) (No EDC)
    playstation | Crash Bandicoot 3 - Warped
    playstation | Crash Bandicoot 2 - Cortex Strikes Back (Europe) (En,Fr,De,Es,It) (EDC)
    playstation | Crash Bandicoot 3 - Buttobi! Sekai Isshuu
```

Alright, there are plenty. Let's pick one of them (remember to use double quotes `""`):

## Download a game

Let's pick `Crash Team Racing`:

```bash
$ ./romie-mac install -t "Crash Team Racing [SCUS-94426]"
```

Output:

```bash
INFO[0000] Using configuration file: /Users/drpaneas/.romie/config.yml
INFO[0000] Installing 1 games ...
INFO[0000] [1/1] - Downloading: "Crash Team Racing [SCUS-94426]"	from	"https://static.emulatorgames.net/roms/playstation/Crash Team Racing [U] [SCUS-94426].rar"
Download progress: 100% - 267/267 MB
INFO[0128] Download completed in 2m8.60368447s
INFO[0128] Extracting the compressed archive ...
INFO[0145] done!
```

The game has been downloaded and extracted as well.
It's available at the default location:

```bash
$ ls -l ~/.romie/downloads/playstation/Crash\ Team\ Racing\ \[SCUS-94426\]/Crash\ Team\ Racing\ \[U\]\ \[SCUS-94426\]
total 1183016
-rw-r--r--  1 drpaneas  staff  605698800 May 17 02:29 CTR - Crash Team Racing (USA).bin
-rw-r--r--  1 drpaneas  staff         95 May 17 02:29 CTR - Crash Team Racing (USA).cue
```

## Delete the game

```
$ ./romie-mac remove -t "Crash Team Racing [SCUS-94426]"
INFO[0000] Using configuration file: /Users/drpaneas/.romie/config.yml
INFO[0000] Removing 1 games ...
INFO[0000] Crash Team Racing [SCUS-94426] has been successfully removed
```

## Change the download location

You can change the directory where the downloaded games will be stored by editing the `config` file as follows:


```bash
$ cat ~/.romie/config.yml
download: "~/.romie/downloads" # <---- Modify this value
database: "~/.romie/database
```
