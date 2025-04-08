#!/bin/sh

# note: the func exists because of this issue:
# https://github.com/deepmap/oapi-codegen/issues/377
fix_external_ref_bug() {
    target="$1"
    import_line="import ("
    ln=$(grep -m1 -n "$import_line" "$target" | cut -d: -f1)
    ed "$target" << END
${ln}a
    externalRef1 "github.com/vpnhouse/api/go/server/common"
.
w
q
END

    sed -i -e 's/*Error/*externalRef1.Error/g' "$target"
    sed -i -e 's/var dest Error/var dest externalRef1.Error/g' "$target"
}

# Loop through all .conf files in the current directory
for conf_name in *.conf; do
    schema_name="../schemas/${conf_name%%.*}.yaml"
    echo "$schema_name..."

    fn="$(basename ${conf_name%.*})"
    gen_type="${fn##*.}"
    if [ $gen_type = "client" ]; then
        mkdir -p "../go/client/${fn%.*}"
    else
        mkdir -p "../go/server/$fn"
    fi

    oapi-codegen -config "$conf_name" "$schema_name"
done

# Define the array of targets using set command
set -- "../go/client/authorizer/authorizer.go" \
    "../go/client/client_service/client_service.go" \
    "../go/client/dashboard/dashboard.go" \
    "../go/client/discovery/discovery.go" \
    "../go/client/federation/federation.go" \
    "../go/client/license_service/license_service.go" \
    "../go/client/project_service/project_service.go" \
    "../go/client/user_service/user_service.go" \
    "../go/client/tunnel_mgmt/tunnel.go" \
    "../go/client/limit_service/limit_service.go"

# Loop through the array of targets
for target; do
    fix_external_ref_bug "$target"
done
