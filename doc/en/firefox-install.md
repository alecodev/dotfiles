## Install Firefox

```bash
sudo su
pacman -S dbus-glib ffmpeg4.4 libxt firejail
cd /
chown alejo:alejo opt/
cd !$
touch /usr/bin/firefox
chmod 755 /usr/bin/firefox
su alejo
```

Get the download URL of the Firefox tar.* file `curl "https://download.mozilla.org/?product=firefox-latest-ssl&os=linux64&lang=en-US"` and change it in the following command
```bash
wget "https://download-installer.cdn.mozilla.net/pub/firefox/releases/98.0.2/linux-x86_64/en-US/firefox-98.0.2.tar.bz2"
tar -xf firefox-*
rm firefox-*
```

Edit file `/usr/bin/firefox` with [text editor][1] and add the following lines
```text
#!/bin/bash
exec firejail /opt/firefox/firefox
```

>Or install with the following command
>```bash
>sudo pacman -S firefox firejail
>sed -i 's/exec \//exec firejail \//' /usr/bin/firefox
>```

Create and change the download directory in Firefox settings to `~/Downloads/Firefox`
```bash
mkdir -p ~/Downloads/Firefox
```

[1]:../../README.md#text-editor