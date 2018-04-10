#!/bin/sh

mkdir -p ~/.ssh
chomd og-rwx ~/.ssh
if [ -n "$SSH_RSA_PRIV_KEY" ]; then
	echo "$SSH_RSA_PRIV_KEY" > ~/.ssh/id_rsa
	chomd og-rwx ~/.ssh/is_rsa
fi
if [ -n "$SSH_KNOWN_HOSTS" ]; then
	echo "$SSH_KNOWN_HOSTS" > ~/.ssh/known_hosts
	chomd og-rwx ~/.ssh/known_hosts
fi
