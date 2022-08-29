## Instalación de Firefox

```bash
sudo su
pacman -S dbus-glib ffmpeg4.4 libxt firejail
cd /
chown alejo:alejo opt/
cd !$
touch /usr/bin/firefox
chmod 755 /usr/bin/firefox
exit
wget --quiet $(curl -s "https://download.mozilla.org/?product=firefox-latest-ssl&os=linux64&lang=en-US"|awk -F'"' '{print $2}')
tar -xf firefox-*
rm firefox-*
```

Edite el archivo `/usr/bin/firefox` con [editor de texto][1] y agregue las siguientes líneas
```text
#!/bin/bash
exec firejail /opt/firefox/firefox
```

>O instala con el siguiente comando
>```bash
>sudo pacman -S firefox firejail
>sed -i 's/exec \//exec firejail \//' /usr/bin/firefox
>```

Cree y cambie el directorio de descarga en la configuración de Firefox a `~/Downloads/Firefox`
```bash
mkdir -p ~/Downloads/Firefox
```

[1]:../../README.es.md#editor-de-texto