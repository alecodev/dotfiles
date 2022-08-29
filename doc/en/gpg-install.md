## Install GPG

```bash
pacman -S gnupg

# Generating a new GPG key
gpg --full-generate-key
git config --global commit.gpgsign true
git config --global credential.helper /usr/lib/git-core/git-credential-libsecret
```

GPG key is added to git
```zsh
gpg --list-secret-keys --keyid-format=long
git config --global user.signingkey YOURKEY

# Generate public key for Github
gpg --armor --export YOURKEY
```

Edit file `~/.bashrc` with [text editor][1] and add the following lines
```zsh
#  GPG key
export GPG_TTY=$(tty)
```

Edit file `~/.zshrc` with [text editor][1] and add the following lines
```zsh
#  GPG key
export GPG_TTY=$TTY
```

[1]:#text-editor