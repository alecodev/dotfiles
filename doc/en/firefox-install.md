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
wget --quiet $(curl -s "https://download.mozilla.org/?product=firefox-latest-ssl&os=linux64&lang=en-US"|awk -F'"' '{print $2}')
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