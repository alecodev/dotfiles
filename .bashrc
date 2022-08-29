#
# ~/.bashrc
#

# If not running interactively, don't do anything
[[ $- != *i* ]] && return

alias ls='ls --color=auto'
PS1='[\u@\h \W]\$ '

# Alias
source ~/.aliases

# Service ssh-agent
export SSH_AUTH_SOCK="$XDG_RUNTIME_DIR/ssh-agent.socket"

#  GPG key
export GPG_TTY=$(tty)

# Editor Default
export EDITOR="/usr/bin/nvim"
export VISUAL="$EDITOR"

# Golang
export PATH=$PATH:/usr/local/go/bin