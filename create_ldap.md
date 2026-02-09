# launch container without authentik-ldap

# Go to the web interface and connect

User is akadmin and you can set the pass with the script

# Go to Admin Interface

# Create a Provider

Lateral Menu -> Application -> Provider
Create :
- Type : LDAP
- Name : ldap-users
- Bind Flow : default-authentification-flow

# Create Application

Lateral Menu -> Applications -> Application
Create :
- Name : LDAP Server
- Slug : ldap-server
- Providers BackChannel : Add ldap-users

# Create a outpost

Lateral Menu -> Applications -> Outposts
Create or Modify
Link the provider ( Type : LDAP )

# Get the existing token

Open the token

# Restart the ldap container

