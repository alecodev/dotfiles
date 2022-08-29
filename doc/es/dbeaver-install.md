## Instalación de DBeaver

```bash
cd /opt
rm -rf dbeaver
wget --quiet https://dbeaver.io/files/dbeaver-ce-latest-linux.gtk.x86_64.tar.gz --output-document=dbeaver-ce.tar.gz
tar -xf dbeaver-ce.tar.gz
rm dbeaver-ce.tar.gz
```

>O instálalo con el siguiente comando
>```bash
>sudo pacman -S dbeaver
>```