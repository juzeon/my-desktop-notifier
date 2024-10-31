# My Desktop Notifier

A simple cross-platform desktop notifier, supporting Windows, Linux and macOS.

## Building

Requirement: Go 1.23+

```bash
go build
# For windows: go build -ldflags -H=windowsgui
./my-desktop-notifier
```

## Configuration

Create `config.yml`:

```yaml
schedules:
  - week: 0 # Sunday
    time: 15:04
    content: Wake up.
  - week: 4 # Thursday
    time: 12:30
    content: Let's go for Crazy Thursday.
  - week: 4
    time: 23:00
    content: Time to sleep.
```

## Windows Setup

Build the app using:

```bash
go build -ldflags -H=windowsgui
```

Then create a shortcut for `my-desktop-notifier.exe` and put it in the startup directory, typically `C:\Users\Administrator\AppData\Roaming\Microsoft\Windows\Start Menu\Programs\Startup`, where `Administrator` should be replaced with your username.

## Screenshots

![image-20241031143739359](https://public.ptree.top/picgo/2024/10/1730356663.png)