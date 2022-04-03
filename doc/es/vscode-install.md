## Instalación de Visual Studio Code

Descarga el archivo tar.gz desde la página `https://code.visualstudio.com/Download` e instálalo con los siguientes comandos
```bash
sudo pacman -S electron
cd /opt
mv ~/Downloads/Firefox/code-*.tar.gz .
tar -xf code-*.tar.gz
rm code-*.tar.gz
cd ~
ln -s /opt/VSCode-linux-x64/bin/code /usr/bin/
```

>O instálalo con yay
>```bash
>yay -S visual-studio-code-bin
>```