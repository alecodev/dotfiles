### Install Google Chrome

```zsh
cd !$
git clone https://aur.archlinux.org/google-chrome.git
cd google-chrome
makepkg -si
cd ..
rm -r google-chrome
sudo su
```

>Or install with the following command
>```zsh
>yay -S google-chrome
>```

Edit file `/usr/bin/google-chrome-stable` with [text editor][1] and add the following lines
```diff
# Launch
-exec /opt/google/chrome/google-chrome $CHROME_USER_FLAGS "$@"
+exec /opt/google/chrome/google-chrome $CHROME_USER_FLAGS "$@" --force-dark-mode
```

Create and change the download directory in Chrome settings to `~/Downloads/Chrome`
```bash
mkdir -p ~/Downloads/Chrome
```

Execute the following commands
```zsh
sudo pacman -S noto-fonts-emoji
gsettings set org.gnome.desktop.interface gtk-theme "Adwaita-dark"
```

[1]:../../README.md#text-editor