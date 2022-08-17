# Dotfiles


***Languages***
- [🇪🇸 - Español](./README.es.md)
- **🇺🇸 - English**
---


#### ***Linux Distribution***
- [Arch](doc/en/arch-install.md)


#### ***Text Editor***
- neovim
>but feel free to use your preferred text editor (vim, nano, ...)


---
## Install Window Manager

Login with the root user and run the following commands
```bash
pacman -Syu
pacman -S gcc make git numlockx nmap wget curl p7zip xclip unzip zsh htop flameshot libsecret gnome-keyring xorg-server xorg-xev bspwm sxhkd alacritty rofi lightdm lightdm-gtk-greeter
```

Edit the file `/etc/lightdm/lightdm.conf` with [text editor][1] and modify the following lines
```diff
-greeter-session=example-gtk-gnome
+greeter-session=lightdm-gtk-greeter
...
-display-setup-script=
+display-setup-script=/usr/bin/setxkbmap -layout latam,es
...
-greeter-setup-script=
+greeter-setup-script=/usr/bin/numlockx on
```

Activate the lightdm service
```bash
systemctl enable lightdm
```

Change user
```bash
su alejo
```

Create or edit the file `~/.xprofile` with [text editor][1] and set the following lines
>**In case of running in a virtual machine like VirtualBox add the following line**
>
>VBoxClient-all &
```text
dbus-update-activation-environment --systemd DISPLAY &
sxhkd &
exec bspwm
```

Create the directories and configuration files of bspwm and sxhkd
```bash
mkdir -p ~/.config/{bspwm,sxhkd}
cd /usr/share/doc/bspwm/examples
cp bspwmrc ~/.config/bspwm
cp sxhkdrc ~/.config/sxhkd
cd ~
chmod +x ~/.config/bspwm/bspwmrc
sed -i 's/urxvt/alacritty/' ~/.config/sxhkd/sxhkdrc
sed -i 's/super + @space/super + r/' ~/.config/sxhkd/sxhkdrc
sed -i 's/dmenu_run/rofi -show run/' ~/.config/sxhkd/sxhkdrc
```

Reboot and log in with the other user
```bash
reboot
```

Ready now you can log in with the other user and use bspwm by pressing `Super + Enter`

Git alias are created
```bash
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status
```

Create the directories and configuration files of bspwm and alacritty
```bash
mkdir -p ~/.config/alacritty
mkdir ~/.config/bspwm/scripts
cd !$
wget --quiet https://raw.githubusercontent.com/alecodev/dotfiles/main/.config/bspwm/scripts/bspwm_resize
wget --quiet https://raw.githubusercontent.com/alecodev/dotfiles/main/.config/bspwm/scripts/bspwm_smart_move
chmod +x ./{bspwm_resize,bspwm_smart_move}
```

Download file `~/.config/sxhkd/sxhkdrc` with the following command
```bash
wget --quiet https://raw.githubusercontent.com/alecodev/dotfiles/main/.config/sxhkd/sxhkdrc --output-document=~/.config/sxhkd/sxhkdrc
```
Press `Super + Alt + r` and  `Super + esc`

Create the bspwm configuration directories in the other user and copy the files (in my case 'alejo')
```bash
touch ~/.aliases
sudo su
ln -sf /home/alejo/.aliases ~/
mkdir ~/.config
ln -s /home/alejo/.config/{bspwm,sxhkd,alacritty} ~/.config/
rm ~/{.xprofile,.bashrc}
ln -s /home/alejo/{.xprofile,.bashrc} ~/
```

---
## Setting up the work environment

### Programs & Applications

- [Firefox](doc/en/firefox-install.md)
- [Fonts](doc/en/fonts-install.md)
- [Yay](doc/en/yay-install.md)
- [Visual Studio Code](doc/en/vscode-install.md)
- [Google Chrome](doc/en/chrome-install.md)

Download file `~/.config/alacritty/alacritty.yml` with the following command
```bash
wget --quiet https://raw.githubusercontent.com/alecodev/dotfiles/main/.config/alacritty/alacritty.yml --output-document=~/.config/alacritty/alacritty.yml
```

### Set Wallpaper
```bash
sudo pacman -S feh
mkdir -p ~/Images
```

Download wallpaper in Images and edit `~/.config/bspwm/bspwmrc` and add next line
```text
feh --bg-fill /home/alejo/Images/wallpaper.jpg
```

### Install Neofetch
```bash
sudo pacman -S neofetch
```

### Install Bluetooth
```bash
sudo pacman -S bluez bluez-utils
```

### Install Polybar
```bash
mkdir -p ~/.config/polybar
sudo su
ln -s /home/alejo/.config/polybar ~/.config/
su alejo
sudo pacman -S cmake pkg-config libuv cairo libxcb python3 xcb-proto xcb-util-image xcb-util-wm python-sphinx python-packaging xcb-util-cursor xcb-util-xrm alsa-lib libpulse i3-wm jsoncpp libmpdclient libnl curl pipewire
cd ~/Downloads/Firefox/

git clone --recursive https://github.com/polybar/polybar
cd polybar

mkdir build
cd build
cmake ..
make -j$(nproc)
sudo make install
cd ../..
rm -r polybar
```

>Or install with the following command
>```bash
>sudo pacman -S polybar
>```

Edit the file `~/.config/bspwm/bspwmrc` with [text editor][1] and add the following line
```bash
polybar &
```

### Install Picom
```bash
sudo pacman -S meson libx11 libxext libconfig libdbus libev pixman uthash xcb-util-image xcb-util-renderutil libgl pcre asciidoc mesa ninja dbus xorg-xprop xorg-xwininfo
cd ~/Downloads/Firefox/

git clone https://github.com/yshui/picom.git
cd picom

git submodule update --init --recursive
meson --buildtype=release . build
ninja -C build
sudo ninja -C build install
cd ..
rm -r picom
```

>Or install with the following command
>```bash
>sudo pacman -S picom
>```

Configure zsh and history
```bash
sudo su
touch ~/{.zshrc,.zsh_history}
su alejo
touch ~/{.zshrc,.zsh_history}
```

Create or edit the file `~/.zshrc` with [text editor][1] and add the following lines
```zsh
# Lines configured by zsh-newuser-install
HISTFILE=~/.zsh_history
HISTSIZE=1000
SAVEHIST=1000
bindkey -e

# End of lines configured by zsh-newuser-install
# The following lines were added by compinstall
zstyle :compinstall filename '/home/alejo/.zshrc'

autoload -Uz compinit
compinit
# End of lines added by compinstall

bindkey "^[[3~" delete-char                     # Key Del
bindkey "^[[5~" beginning-of-buffer-or-history  # Key Page Up
bindkey "^[[6~" end-of-buffer-or-history        # Key Page Down
bindkey "^[[H" beginning-of-line                # Key Home
bindkey "^[[F" end-of-line                      # Key End
bindkey "^[[1;3C" forward-word                  # Key Alt + Right
bindkey "^[[1;3D" backward-word                 # Key Alt + Left
```

### Install Powerlevel10k
```bash
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
echo 'source ~/powerlevel10k/powerlevel10k.zsh-theme' >>~/.zshrc
zsh

sudo su
usermod --shell /usr/bin/zsh alejo
usermod --shell /usr/bin/zsh root

git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
ln -sf /home/alejo/.zshrc ~/
ln -sf /home/alejo/.p10k.zsh ~/
zsh

su alejo
```

Edit file `~/.zshrc` and `~/.bashrc` with [text editor][1] and add the following lines
```bash
# Alias
source ~/.aliases
```

### Install Docker
```zsh
sudo su
pacman -S docker docker-compose
systemctl enable docker
systemctl restart docker
groupadd -r -g 82 www-data
useradd -M -r -u 82 -g 82 -c "User HTTP files" -s /usr/bin/nologin www-data
usermod -aG docker,www-data alejo
```

### Install DBeaver
```zsh
cd /opt
rm -rf dbeaver
wget --quiet https://dbeaver.io/files/dbeaver-ce-latest-linux.gtk.x86_64.tar.gz --output-document=dbeaver-ce.tar.gz
tar -xf dbeaver-ce.tar.gz
rm dbeaver-ce.tar.gz
```

>Or install with the following command
>```zsh
>sudo pacman -S dbeaver
>```

Download file `~/.aliases` with the following command
```bash
wget --quiet https://raw.githubusercontent.com/alecodev/dotfiles/main/.aliases --output-document=~/.aliases
```

### Install bat, lsd & fzf
```zsh
sudo pacman -S bat lsd fzf
```

### Install zsh plugins (sudo, autosuggestions, syntax highlighting)
```zsh
sudo su
mkdir -p /usr/share/zsh/plugins
cd !$
chown alejo:alejo /usr/share/zsh/plugins
su alejo
mkdir sudo
cd !$
wget --quiet https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/plugins/sudo/sudo.plugin.zsh
sudo pacman -S zsh-autosuggestions zsh-syntax-highlighting
```

Edit file `~/.zshrc` with [text editor][1] and add the following lines
```zsh
# Plugins
source /usr/share/zsh/plugins/sudo/sudo.plugin.zsh
source /usr/share/zsh/plugins/zsh-autosuggestions/zsh-autosuggestions.zsh
source /usr/share/zsh/plugins/zsh-syntax-highlighting/zsh-syntax-highlighting.zsh
```

### Install SSH
```zsh
sudo pacman -S openssh
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
mkdir -p ~/.config/systemd/user
mkdir -p ~/.ssh
chmod 700 ~/.ssh
touch ~/.ssh/config
chmod 600 ~/.ssh/config
echo "# host-specific options\n" >> ~/.ssh/config
```

Download file `~/.config/systemd/user/ssh-agent.service` with the following command
```bash
wget --quiet https://raw.githubusercontent.com/alecodev/dotfiles/main/.config/systemd/user/ssh-agent.service --output-document=~/.config/systemd/user/ssh-agent.service
```

Edit file `~/.bashrc` and `~/.zshrc` with [text editor][1] and add the following lines
```zsh
# Service ssh-agent
export SSH_AUTH_SOCK="$XDG_RUNTIME_DIR/ssh-agent.socket"
```

Enable the service
```zsh
systemctl --user enable ssh-agent
systemctl --user start ssh-agent
ssh-add ~/.ssh/id_rsa
```

### Generating a GPG key
```zsh
gpg --full-generate-key
git config --global commit.gpgsign true
git config --global credential.helper /usr/lib/git-core/git-credential-libsecret
```

GPG key is added to git
```zsh
gpg --list-secret-keys --keyid-format=long
git config --global user.signingkey YOURKEY
```

Edit file `~/.bashrc` with [text editor][1] and add the following lines
```zsh
#  GPG key
export GPG_TTY=$(tty)

# Editor Default
export EDITOR="/usr/bin/nvim"
export VISUAL="$EDITOR"
```

Edit file `~/.zshrc` with [text editor][1] and add the following lines
```zsh
#  GPG key
export GPG_TTY=$TTY

# Editor Default
export EDITOR="/usr/bin/nvim"
export VISUAL="$EDITOR"
```

[1]:#text-editor