## Instalación de Yay

```bash
cd /tmp
git clone https://aur.archlinux.org/yay.git
cd yay
makepkg -si
cd ..
rm -r yay
yay -Yc
```