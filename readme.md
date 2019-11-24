![alt](/img/gimpy_logo.png)
# Gopher Imp - YAML based zDoom launcher

a Go port of Python [yzdoom](https://github.com/squeaky-godzilla/yzdoom) with some minor changes, aimed at cross-platform compatibility and minimized overhead.

## What it does
**gimpy** allows you to create a [zdoom](https://zdoom.org/index) run setup in YAML, which means minimal overhead no-nonsense game launch with a trendy devops feeling to it.

## How does it work
much like [yzdoom](https://github.com/squeaky-godzilla/yzdoom), **gimpy** takes a command line argument of a *DoomRun* config, where you can specify what iWAD and what pWADs you want to use & location of your *DoomDefault* config.

DoomRun example:

```
DoomRun:
  defaultsPath: ./defaults.yml
  iwad: doom2.wad
  pwads: 
    - brutalv21.pk3
```

DoomDefault example:

```
DoomDefault:
  gzdoom: /usr/games/gzdoom
  iwadFolder: /home/vitek/.config/gzdoom
  pwadFolder: /home/vitek/.config/gzdoom
```

