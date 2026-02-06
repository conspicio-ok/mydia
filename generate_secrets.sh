#!/bin/bash

echo "=== Génération des secrets pour Authentik ==="
echo ""

# PostgreSQL Password
echo "📦 PostgreSQL Password:"
PG_PASS=$(openssl rand -base64 32)
echo "PG_PASS=$PG_PASS"
echo ""

# Authentik Secret Key
echo "🔑 Authentik Secret Key:"
AUTHENTIK_SECRET_KEY=$(openssl rand -base64 60)
echo "AUTHENTIK_SECRET_KEY=$AUTHENTIK_SECRET_KEY"
echo ""

# LDAP Bind Password (pour le compte service)
echo "🔐 LDAP Bind Password (pour ldap-service user):"
LDAP_BIND_PASSWORD=$(openssl rand -base64 32)
echo "LDAP_BIND_PASSWORD=$LDAP_BIND_PASSWORD"
echo ""

echo "=== Création du fichier .env ==="
cat > .env << EOF
# === PostgreSQL ===
PG_PASS=$PG_PASS

# === Authentik Secret Key ===
AUTHENTIK_SECRET_KEY=$AUTHENTIK_SECRET_KEY

# === LDAP Token ===
# À générer après l'installation d'Authentik
# 1. Se connecter à http://auth.homeserver.local
# 2. Admin Interface → Tokens → Create Token
# 3. Copier le token ici
AUTHENTIK_LDAP_TOKEN=

# === LDAP Bind Password ===
LDAP_BIND_PASSWORD=$LDAP_BIND_PASSWORD

# === Email (Optionnel) ===
# Décommenter si tu veux envoyer des emails
# EMAIL_PASSWORD=
EOF

echo "✅ Fichier .env créé avec succès !"
echo ""
echo "⚠️  N'oublie pas de générer le AUTHENTIK_LDAP_TOKEN après l'installation"
echo "    et de l'ajouter dans le fichier .env"
echo ""
echo "📋 Sauvegarde ces mots de passe dans un gestionnaire de mots de passe :"
echo "    - PG_PASS (PostgreSQL)"
echo "    - AUTHENTIK_SECRET_KEY"
echo "    - LDAP_BIND_PASSWORD (pour créer le user ldap-service)"
