## Install Visual Studio Code

Download file tar.gz from the page `https://code.visualstudio.com/Download` and install with the following commands
```bash
sudo pacman -S electron
cd /opt
mv ~/Downloads/Firefox/code-*.tar.gz .
tar -xf code-*.tar.gz
rm code-*.tar.gz
cd ~
ln -s /opt/VSCode-linux-x64/bin/code /usr/bin/
```

>Or install with yay
>```bash
>yay -S visual-studio-code-bin
>```