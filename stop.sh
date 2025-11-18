#!/bin/bash
docker-compose stop
sleep 5
docker-compose down
apps_dir="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)/apps"
if [ -d "$apps_dir" ]; then
    find "$apps_dir" -maxdepth 1 -type f \
        ! -name 'client-mtls-api.json' \
        ! -name 'keyless-plugin-api.json' \
        ! -name 'protected-api.json' \
        -delete
fi
rm -rf tyk_apis