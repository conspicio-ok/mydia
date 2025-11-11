# Mydia

Mydia an ensemble of services to be independent.

## Services :

- VPN : wg-easy
- Monitoring : traefik
- SSO : authelia
- Homepage : homarr
- Music : navidrome
- Streaming : jellyfin

## Unvisible services :

- DNS : dnsmasq
- Certificats : mkcert-generator
- Network (bridge : 172.20.0.0/24)
- : go-api
- : redis

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

To launch the services, run :
```bash
docker compose up -d
```

and go to
`http://YOUR_IP:51821/`

## Warning

### Wg-easy

Maybe you have a bug (10.13.13.0) for the ip client, modify it to 10.13.13.2 and increment the next.
Change the /24 to /32 !
In interface section add `WG_MTU=1420`

## Authelia

Default password is `changeme`
So PLEASE change it !
But change equaly the name, display username, and the username