#!/bin/sh

mkdir -p ~/.ssh
chmod og-rwx ~/.ssh
if [ -n "$SSH_RSA_PRIV_KEY" ]; then
	echo "$SSH_RSA_PRIV_KEY" > ~/.ssh/id_rsa
	chmod og-rwx ~/.ssh/id_rsa
fi
if [ -n "$SSH_KNOWN_HOSTS" ]; then
	echo "$SSH_KNOWN_HOSTS" > ~/.ssh/known_hosts
	chmod og-rwx ~/.ssh/known_hosts
fi
