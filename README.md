# Dotfiles in Arch Linux

***Languages***
- [🇪🇸 - Español](./README.es.md)
- **🇺🇸 - English**


#### ***Text Editors***
- neovim
- vim
- nano


## Arch Linux installation
For more information on the installation process visit [Arch's guide](https://wiki.archlinux.org/title/Installation_guide)

>If you are installing on a virtual machine
>- Enable EFI
>- Disable 3D acceleration


Sets the console keyboard layout (in my case 'es'). You can validate the available keyboard layouts with the following command: `ls /usr/share/kbd/keymaps/**/*.map.gz`
```bash
loadkeys es
```

Verify the boot mode EFI
```bash
ls /sys/firmware/efi/efivars
```

Check connection
```bash
ip link
ping 8.8.8.8
```

Update the system clock
```bash
timedatectl set-ntp true
```

Identifies the device block usually is /dev/sda
```bash
lsblk
```

Set the partitions
```bash
cfdisk /dev/sda
```

Select GPT and create the following partitions (the size depends on the use you want to give it), remember to write before exiting
```
512M      EFI System         (This will be the size of the system boot partition)
16G       Linux Swap         (This will be the size of the SWAP memory, it is recommended to double the size of the RAM memory)
40G       Linux filesystem   (This will be the size assigned to /)
63.5G     Linux filesystem   (This will be the size assigned to /home)
write
```

Check the partitions
```bash
lsblk
```

Set the format of the partitions
```bash
mkfs.fat -F32 /dev/sda1
mkswap /dev/sda2
mkfs.ext4 /dev/sda3
mkfs.ext4 /dev/sda4
```

Mount the partitions
```bash
swapon /dev/sda2
mount /dev/sda3 /mnt
mkdir /mnt/{efi,home}
mount /dev/sda1 /mnt/efi
mount /dev/sda4 /mnt/home
```

Check the partitions
```bash
lsblk
```

Install the basic packages
```bash
pacstrap /mnt base base-devel linux linux-firmware neovim dhcpcd
```

Generate the Fstab file
```bash
genfstab -U /mnt >> /mnt/etc/fstab
```

Change root into the new system
```bash
arch-chroot /mnt
```

Set the time zone (in my case 'America/Bogota'), you can see the available time zones with the following command: `timedatectl list-timezones`
```bash
ln -sf /usr/share/zoneinfo/America/Bogota /etc/localtime
hwclock --systohc
```

Set Localization, edit the file `/etc/locale.gen` and uncomment `en_US.UTF-8 UTF-8` with your preferred [text editor](#text-editors)

Generate the regional settings and set the default keyboard layout by running
```bash
locale-gen
echo "LANG=en_US.UTF-8" >> /etc/locale.conf
echo "KEYMAP=es" >> /etc/vconsole.conf
```

Create the hostname file (in my case the hostname will be 'Arch')
```bash
echo "Arch" >> /etc/hostname
```

Add the domains to the file `/etc/hosts` with your preferred [text editor](#text-editors), replace the name of the computer with the one you established in the previous step
```bash
127.0.0.1   localhost
::1         localhost
127.0.1.1   Arch.localdomain  Arch
```

Activate the DHCPCD service
```bash
systemctl enable dhcpcd.service
```

Set the password for the root user
```bash
passwd
```

Create a new user (in my case 'alejo')
```bash
useradd -m alejo
passwd alejo
usermod -aG wheel,audio,video,optical,storage alejo
```

Activate the sudo group by executing the following command and uncomment `%wheel ALL=(ALL) ALL`
```bash
EDITOR=nvim
visudo
```


>In case of running in a virtual machine like VirtualBox run the following command
>```bash
>pacman -S virtualbox-guest-utils
>systemctl enable vboxservice
>usermod -aG vboxsf alejo
>```

Set the bootloader
```bash
pacman -S grub sudo efibootmgr os-prober intel-ucode
grub-install --target=x86_64-efi --efi-directory=/efi --bootloader-id=GRUB
echo "GRUB_DISABLE_OS_PROBER=false" >> /etc/default/grub
grub-mkconfig -o /boot/grub/grub.cfg
```

Finish the process with the following command
```bash
exit
umount -R /mnt
reboot
```

---
## Install Window Manager

Login with the root user and run the following commands
```bash
pacman -Syu
pacman -S gcc make git libsecret xorg-server bspwm sxhkd alacritty rofi lightdm lightdm-gtk-greeter numlockx zsh neovim htop xorg-xev nmap
```

Edit file `/etc/pacman.conf` with your preferred [text editor](#text-editors) and modify the following lines
```diff
# Misc options
#UseSyslog
-#Color
+Color
#NoProgressBar
CheckSpace
#VerbosePkgLists
+ILoveCandy
```

Edit the file `/etc/lightdm/lightdm.conf` with your preferred [text editor](#text-editors) and modify the following lines
```diff
-greeter-session=example-gtk-gnome
+greeter-session=lightdm-gtk-greeter
...
-display-setup-script=
+display-setup-script=/usr/bin/setxkbmap es
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

Create or edit the file `~/.xprofile` with your preferred [text editor](#text-editors) and set the following lines
```text
dbus-update-activation-environment --systemd DISPLAY &
VBoxClient-all &
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
```

Reboot and log in with the other user
```bash
reboot
```

Ready now you can log in with the other user and use bspwm by pressing `Super + Enter`

Git aliases are created
```bash
git config --global alias.co checkout
git config --global alias.br branch
git config --global alias.ci commit
git config --global alias.st status
```

Configure alias, zsh and history
```bash
sudo su
touch ~/{.zshrc,.zsh_history}
su alejo
touch ~/{.aliases,.zshrc,.zsh_history}
```

Create or edit the file `~/.aliases` with your preferred [text editor](#text-editors) and add the following lines
```bash
# ls
alias ls='lsd --group-dirs=first'
alias ll='lsd -lh --group-dirs=first'
alias la='lsd -Alh --group-dirs=first'
alias l='lsd --group-dirs=first'

# cat
alias cat='/bin/bat'
alias catn='/bin/cat'
alias catnl='/bin/bat --paging=never'

# diff
alias diff='/usr/bin/diff --color=auto'

# cd
alias cd..='cd ..'

# Colorize the grep command output for ease of use (good for log files)
alias grep='grep --color=auto'
alias egrep='egrep --color=auto'
alias fgrep='fgrep --color=auto'
```

Create or edit the file `~/.bashrc` with your preferred [text editor](#text-editors) and add the following lines
```bash
# Alias
source ~/.aliases
```

Create or edit the file `~/.zshrc` with your preferred [text editor](#text-editors) and add the following lines
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

Create the directories and configuration files of bspwm, polybar and alacritty
```bash
mkdir -p ~/.config/{polybar,alacritty}
mkdir ~/.config/bspwm/scripts
touch ~/.config/bspwm/scripts/{bspwm_resize,bspwm_smart_move}
chmod +x ~/.config/bspwm/scripts/{bspwm_resize,bspwm_smart_move}
```

Edit file `~/.config/bspwm/scripts/bspwm_resize` with your preferred [text editor](#text-editors) and add the following lines
```bash
#!/usr/bin/env bash

if bspc query -N -n focused.floating > /dev/null; then
	step=20
else
	step=100
fi

case "$1" in
	west) dir=right; falldir=left; x="-$step"; y=0;;
	east) dir=right; falldir=left; x="$step"; y=0;;
	north) dir=top; falldir=bottom; x=0; y="-$step";;
	south) dir=top; falldir=bottom; x=0; y="$step";;
esac

bspc node -z "$dir" "$x" "$y" || bspc node -z "$falldir" "$x" "$y"
```

Edit file `~/.config/bspwm/scripts/bspwm_smart_move` with your preferred [text editor](#text-editors) and add the following lines
```bash
#!/bin/bash

# based on https://gitlab.com/protesilaos/dotfiles/-/blob/v2.2.0/bin/bin/bspwm_smart_move

[ "$#" -eq 1 ] || { echo "Pass only one argument: north,east,south,west"; exit 1; }

# Check if argument is a valid direction.
case "$1" in
    north|east|south|west)
        dir="$1"
        ;;
    *)
        echo "Not a valid argument."
        echo "Use one of: north,east,south,west"
        exit 1
        ;;
esac

_query_nodes() {
    bspc query -N -n "$@"
}

# Do not operate on floating windows!
#[ -z "$(_query_nodes focused.floating)" ] || { echo "Only move tiled windows."; exit 1; }

receptacle="$(_query_nodes 'any.leaf.!window')"

# This regulates the behaviour documented in the description.
if [ -n "$(_query_nodes focused.floating)" ]; then
    case "$1" in
        west) x=-20; y=0;;
        south) x=0; y=20;;
        north) x=0; y=-20;;
        east) x=20; y=0;;
    esac
    bspc node -v "$x" "$y"
elif [ -n "$(_query_nodes "${dir}.!floating")" ]; then
    bspc node -s "$dir"
elif [ -n "$receptacle" ]; then
    bspc node focused -n "$receptacle" --follow
else
    bspc node @/ -p "$dir" -i && bspc node -n "$receptacle" --follow
fi
```

Edit file `~/.config/sxhkd/sxhkdrc` with your preferred [text editor](#text-editors) and modify the following lines
```diff
# terminal emulator
super + Return
	alacritty

# program launcher
-super + @space
-	dmenu_run
+super + r
+	rofi -show run

...

# focus the node in the given direction
-super + {_,shift + }{h,j,k,l}
+super + {_,shift + }{Left,Down,Up,Right}

...

# focus the last node/desktop
-super + {grave,Tab}
-	bspc {node,desktop} -f last
+#super + {grave,Tab}
+#	bspc {node,desktop} -f last
+
+# focus on the next or previous desktop
+super + {shift + ,_} Tab
+	bspc desktop -f {prev,next}.local.occupied

# focus the older or newer node in the focus history
super + {o,i}
	bspc wm -h off; \
	bspc node {older,newer} -f; \
	bspc wm -h on

# focus or send to the given desktop
super + {_,shift + }{1-9,0}
	bspc {desktop -f,node -d} '^{1-9,10}'

+# Custom numpad focus or send to the given desktop
+super + {_,shift + }KP_{1-9,End,Down,Next,Left,Begin,Right,Home,Up,Prior,0,Insert}
+        bspc {desktop -f,node -d} '^{1-9,1,2,3,4,5,6,7,8,9,10,10}'

#
# preselect
#

# preselect the direction
-super + ctrl + {h,j,k,l}
+super + ctrl + alt + {Left,Down,Up,Right}

...

# cancel the preselection for the focused desktop
-super + ctrl + shift + space
+super + ctrl + alt + space

...

# expand a window by moving one of its side outward
-super + alt + {h,j,k,l}
-	bspc node -z {left -20 0,bottom 0 20,top 0 -20,right 20 0}
+#super + alt + {h,j,k,l}
+#	bspc node -z {left -20 0,bottom 0 20,top 0 -20,right 20 0}

# contract a window by moving one of its side inward
-super + alt + shift + {h,j,k,l}
-	bspc node -z {right -20 0,top 0 20,bottom 0 -20,left 20 0}
+#super + alt + shift + {h,j,k,l}
+#	bspc node -z {right -20 0,top 0 20,bottom 0 -20,left 20 0}

# move a floating window
-super + {Left,Down,Up,Right}
-       bspc node -v {-20 0,0 20,0 -20,20 0}
+#super + {Left,Down,Up,Right}
+#       bspc node -v {-20 0,0 20,0 -20,20 0}

+# Custom move window
+super + ctrl + {Left,Down,Up,Right}
+        /home/alejo/.config/bspwm/scripts/bspwm_smart_move {west,south,north,east}

+# Custom resize window
+alt + super + {Left,Down,Up,Right}
+	/home/alejo/.config/bspwm/scripts/bspwm_resize {west,south,north,east}
```

Create the bspwm configuration directories in the other user and copy the files (in my case 'alejo')
```bash
sudo su
ln -sf /home/alejo/.aliases ~/
ln -s /home/alejo/.config/{bspwm,sxhkd,polybar,alacritty} ~/.config/
rm ~/.xprofile
ln -s /home/alejo/.xprofile ~/
rm ~/.bashrc
ln -s /home/alejo/.bashrc ~/
```

---
## Setting up the work environment

### Install Firefox
```bash
sudo su
pacman -Syu
pacman -S wget curl libstdc++5 dbus-glib unzip p7zip firejail
cd /
chown alejo:alejo opt/
cd !$
touch /usr/bin/firefox
chmod 755 /usr/bin/firefox
su alejo
```

Get the download URL of the Firefox tar.* file `curl "https://download.mozilla.org/?product=firefox-latest-ssl&os=linux64&lang=en-US"` and change it in the following command
```bash
wget "https://download-installer.cdn.mozilla.net/pub/firefox/releases/95.0.2/linux-x86_64/en-US/firefox-95.0.2.tar.bz2"
tar -xf firefox-*
rm firefox-*
mkdir -p ~/Downloads/Firefox
```

Edit file `/usr/bin/firefox` with your preferred [text editor](#text-editors) and add the following lines
```text
#!/bin/bash
exec firejail /opt/firefox/firefox
```

Change the download directory in Firefox settings to `~/Downloads/Firefox`

Edit file `~/.config/sxhkd/sxhkdrc` with your preferred [text editor](#text-editors) and add the following lines
```text
# Open Firefox
super + shift + f
        /usr/bin/firefox
```

Press `Super + Alt + r` and  `Super + esc`, open Firefox with `Super + Shift + f`

### Install Fonts
```bash
sudo su
```

- #### ***Hack Nerd Font***
	---
	Download the font ***Hack Nerd Font*** from the page `https://www.nerdfonts.com/font-downloads` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/Hack
	cd !$
	mv /home/alejo/Downloads/Firefox/Hack.zip .
	unzip Hack.zip
	rm Hack.zip
	```

- #### ***Meslo Nerd Font***
	---
	Download the font ***Meslo Nerd Font*** from the page `https://www.nerdfonts.com/font-downloads` and install with the following commands
	```bash
	mkdir -p /usr/local/share/fonts/nerd-fonts/Meslo
	cd !$
	mv /home/alejo/Downloads/Firefox/Meslo.zip .
	unzip Meslo.zip
	rm Meslo.zip
	```

Reload fonts
```bash
fc-cache -vf
```

Copy the configuration file from alacritty
```bash
su alejo
cp /usr/share/doc/alacritty/example/alacritty.yml ~/.config/alacritty/
```

Edit file `~/.config/alacritty/alacritty.yml` with your preferred [text editor](#text-editors) and modify the following lines
```diff
# Font configuration
-#font:
+font:
  # Normal (roman) font face
-  #normal:
+  normal:
    # Font family
    #
    # Default:
    #   - (macOS) Menlo
    #   - (Linux/BSD) monospace
    #   - (Windows) Consolas
-    #family: monospace
+    family: "Hack Nerd Font Mono"

    # The `style` can be specified to pick a specific face.
-    #style: Regular
+    style: Regular

  # Bold font face
-  #bold:
+  bold:
    # Font family
    #
    # If the bold family is not specified, it will fall back to the
    # value specified for the normal font.
-    #family: monospace
+    family: "Hack Nerd Font Mono"

    # The `style` can be specified to pick a specific face.
-    #style: Bold
+    style: Bold

  # Italic font face
-  #italic:
+  italic:
    # Font family
    #
    # If the italic family is not specified, it will fall back to the
    # value specified for the normal font.
-    #family: monospace
+    family: "Hack Nerd Font Mono"

    # The `style` can be specified to pick a specific face.
-    #style: Italic
+    style: Italic

  # Bold italic font face
-  #bold_italic:
+  bold_italic:
    # Font family
    #
    # If the bold italic family is not specified, it will fall back to the
    # value specified for the normal font.
-    #family: monospace
+    family: "Hack Nerd Font Mono"

    # The `style` can be specified to pick a specific face.
-    #style: Bold Italic
+    style: Bold Italic

  # Point size
-  #size: 11.0
+  size: 12
```

### Install Visual Studio Code
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

### Set Wallpaper
```bash
sudo pacman -S feh
mkdir -p ~/Desktop/alejo/Images
```

Download wallpaper in Images and edit `~/.config/bspwm/bspwmrc` and add next line
```text
feh --bg-fill /home/alejo/Desktop/alejo/Images/wallpaper.jpg
```

### Install Neofetch
```bash
sudo pacman -S neofetch
```

### Install Polybar
```bash
sudo pacman -S cmake pkg-config libuv cairo libxcb python3 xcb-proto xcb-util-image xcb-util-wm python-sphinx python-packaging xcb-util-cursor xcb-util-xrm alsa-lib libpulse i3-wm jsoncpp libmpdclient libnl curl pulseaudio
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

### Install Powerlevel10k
```bash
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
echo 'source ~/powerlevel10k/powerlevel10k.zsh-theme' >>~/.zshrc
zsh

sudo su
git clone --depth=1 https://github.com/romkatv/powerlevel10k.git ~/powerlevel10k
echo 'source ~/powerlevel10k/powerlevel10k.zsh-theme' >>~/.zshrc
zsh

ln -sf /home/alejo/.zshrc ~/

usermod --shell /usr/bin/zsh alejo
usermod --shell /usr/bin/zsh root
su alejo
```

### Install bat
```zsh
sudo pacman -S bat
```

Edit file `~/.zshrc` with your preferred [text editor](#text-editors) and add the following lines
```zsh
# Alias
source ~/.aliases
```

### Install lsd
```zsh
sudo pacman -S lsd
```

### Install fzf
```zsh
git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
~/.fzf/install
sudo su
git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
~/.fzf/install
```

### Install sudo zsh plugin
```zsh
sudo su
mkdir -p /usr/share/zsh-plugins
cd !$
chown alejo:alejo /usr/share/zsh-plugins
su alejo
wget https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/plugins/sudo/sudo.plugin.zsh
```

Edit file `~/.zshrc` with your preferred [text editor](#text-editors) and add the following lines
```zsh
source /usr/share/zsh-plugins/sudo.plugin.zsh
```

### Install SSH
```zsh
sudo pacman -S openssh
ssh-keygen -t rsa -b 4096 -C "your_email@example.com"
mkdir -p ~/.config/systemd/user/
touch ~/.config/systemd/user/ssh-agent.service
touch ~/.ssh/config
chmod 600 .ssh/config
echo "# host-specific options\n" >> ~/.ssh/config
```

Edit file `~/.config/systemd/user/ssh-agent.service` with your preferred [text editor](#text-editors) and add the following lines
```zsh
[Unit]
Description=SSH key agent

[Service]
Type=simple
Environment=SSH_AUTH_SOCK=%t/ssh-agent.socket
# DISPLAY required for ssh-askpass to work
Environment=DISPLAY=:0
ExecStart=/usr/bin/ssh-agent -D -a $SSH_AUTH_SOCK

[Install]
WantedBy=default.target
```

Edit file `~/.bashrc` and `~/.zshrc` with your preferred [text editor](#text-editors) and add the following lines
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

### Install Google Chrome
```zsh
sudo pacman -S noto-fonts-emoji
mkdir -p ~/Downloads/Chrome
cd !$
git clone https://aur.archlinux.org/google-chrome.git
cd google-chrome
makepkg -si
cd ..
rm -r google-chrome
sudo su
```

Edit file `/usr/bin/google-chrome-stable` with your preferred [text editor](#text-editors) and add the following lines
```diff
# Launch
-exec /opt/google/chrome/google-chrome $CHROME_USER_FLAGS "$@"
+exec firejail /opt/google/chrome/google-chrome $CHROME_USER_FLAGS "$@" --force-dark-mode
```

Change the download directory in Chrome settings to `~/Downloads/Chrome`

Edit file `~/.config/sxhkd/sxhkdrc` with your preferred [text editor](#text-editors) and add the following lines
```text
# Open Chrome
super + shift + g
        /usr/bin/google-chrome-stable
```

# Install Mysql Workbench
```zsh
sudo pacman -S mysql-workbench gnome-keyring
dbus-update-activation-environment --systemd DISPLAY
sudo su
cd /usr/share/mysql-workbench/data
mv code_editor.xml code_editor.xml_original
wget https://raw.githubusercontent.com/mleandrojr/mysql-workbench-dark-theme/master/code_editor.xml
```

# Generating a GPG key
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

Edit file `~/.bashrc` with your preferred [text editor](#text-editors) and add the following lines
```zsh
#  GPG key
export GPG_TTY=$(tty)

# Editor Default
export EDITOR="/usr/bin/nvim"
export VISUAL="$EDITOR"
```

Edit file `~/.zshrc` with your preferred [text editor](#text-editors) and add the following lines
```zsh
#  GPG key
export GPG_TTY=$TTY

# Editor Default
export EDITOR="/usr/bin/nvim"
export VISUAL="$EDITOR"
```

# Install Flameshot
```zsh
sudo pacman -S flameshot
```

Edit file `~/.config/sxhkd/sxhkdrc` with your preferred [text editor](#text-editors) and add the following lines
```text
# Open Flameshot
super + shift + s
        /usr/bin/flameshot gui
```

# Install Docker
```zsh
sudo su
pacman -S docker docker-compose
systemctl enable docker
systemctl restart docker
groupadd -r -g 82 www-data
useradd -M -r -u 82 -g 82 -c "User HTTP files" -s /usr/bin/nologin www-data
usermod -aG docker,www-data alejo
```

Edit file `~/.aliases` with your preferred [text editor](#text-editors) and add the following lines
```bash
# Docker
alias dps="docker ps"
alias dpsa="docker ps -a"
alias dim="docker images"
alias dc="docker-compose"
alias dcu="docker-compose up -d"
alias dcd="docker-compose down"
alias dcl="docker-compose logs"
```