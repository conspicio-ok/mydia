# Mydia

Mydia an ensemble of services to be independent.

## Services :

- VPN : wg-easy
- Monitoring : traefik
- SSO : authelia + redis
- Homepage : homarr
- Music : navidrome
- Streaming : jellyfin

### Unvisible services :

- DNS : dnsmasq
- Certificats : mkcert-generator
- Network (bridge : 172.20.0.0/24)

---

## Install

To get the repository, run :
```bash
git clone https://github.com/conspicio-ok/mydia.git
```

You can put all files in a specific directory with
```bash
mkdir /app
mv ./* /app
```

## Before first launch

*If you have already run `docker compose up -d` you can only run*
```bash
docker compose down
```

Change the `authelia/users_database.yml` with your username, display username and email !

Rename the env to `.env`
```bash
mv env .env
```

and modify it
```bash
vim .env
```
or
```bash
nano .env
```

Finally,
To launch the services, run :
```bash
docker compose up -d
```

and go to
`http://YOUR_IP_SERVER:51821/`


## Warning

### Wg-easy

Maybe you have a bug (10.13.13.0) for the ip client, modify it to 10.13.13.2 and increment the next.
Change the /24 to /32 !
In interface section add `WG_MTU=1420`
For the moment, bug is detected on MacOs, if you connect for the first time to the vpn, or after restart container.
Please connect to the wg-page, run
```bash
nslookup google.com
```
for active the dns

### Authelia

Default password is `changeme`
So PLEASE change it after your first connection !
But change equaly the name, display username, and the username