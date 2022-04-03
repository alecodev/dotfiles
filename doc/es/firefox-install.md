## Instalación de Firefox

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

Obtén la URL de descarga del archivo tar.* de Firefox `curl "https://download.mozilla.org/?product=firefox-latest-ssl&os=linux64&lang=en-US"` y cámbiala en el siguiente comando
```bash
wget "https://download-installer.cdn.mozilla.net/pub/firefox/releases/98.0.2/linux-x86_64/en-US/firefox-98.0.2.tar.bz2"
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