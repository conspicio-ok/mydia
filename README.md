# Mydia

Mydia an ensemble of services to be independent.

## Services :

- VPN : wg-easy
- Monitoring : traefik
- SSO : authelia
- : redis
- Homepage : homarr
- Music : navidrome
- Streaming : jellyfin

## Unvisible services :

- DNS : dnsmasq
- Certificats : mkcert-generator
- Network (bridge : 172.20.0.0/24)
- go-api

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
