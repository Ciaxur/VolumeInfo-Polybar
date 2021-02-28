# VolumeInfo-Polybar
> Displays the current Sink's Volume Information for Polybar

## Build 🔨
Use the `go build` Command to compile the application.
``` bash
go build -o volume-info ./src
```

## Usage 🚀
There are two optional Arguments to Change the Color Output of the Icon and String.

**WARNING**: `pacmd` is required to run this application. Go to [Arch Wiki's PulseAudio Page](https://wiki.archlinux.org/index.php/PulseAudio).

``` bash
# Defaults Color to White
volume-info

# Arg1 = Icon Color
# Arg2 = String Color
volume-info "#FF" "#22"
```

## License
This project is licensed under the [MIT License](LICENSE).