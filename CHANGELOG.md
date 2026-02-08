# Changelog

Toutes les modifications notables de ce projet seront documentées dans ce fichier.

Le format est basé sur [Keep a Changelog](https://keepachangelog.com/fr/1.0.0/),
et ce projet adhère au [Semantic Versioning](https://semver.org/lang/fr/).

Les modifications seront repartie selon trois categories :
- **Added**
- **Changed**
- **Removed**

Ces modifications seront alors sous une baniere : **[X.Y.Z] - YYYY-MM-DD**

Les informations a transmettre d'un version a l'autre sont en **gras**,
Toutes **informations** qui n'a pas ete traite doit etre transmise a nouveau


## [0.0.2] - 2026

### Added
- create a docker compose for authentik (LDAP system)

### Changed
- Adapt the other services to use the SSO

ip authentik :
```text
http://172.18.0.5:9000/
```

## [0.0.1] - 2026-02-04

### Added
- create a docker compose for traefik

### Changed
- Add label for navidrome can connect with traefik

```bash
curl -H "Host: navidrome.homeserver.local" http://localhost
```
```bash
curl -H "Host: traefik.homeserver.local" http://localhost:8080
```
Pas de dns pour l'instant les ip sont :
```text
localhost:8080
```
```text
172.18.0.3:4533
```

## [0.0.0] - 2026-02-04

### Added
- create a backup and a navidrome docker compose

Starting of the project
