Collection of commands and Ansible Playbooks for migrating from VM to Kubernetes

1. App with DB
```sh
vagrant up --provision
vagrant ssh appsrv0 -c "curl appsrv0:8080"
```
2. Add a second app
```sh
cat <<EOF >> plays/group_vars/appsrv.yaml
app_names:
- "app"
- "app2"
EOF

vagrant provision
vagrant ssh appsrv0 -c "curl appsrv0:8081"
```
