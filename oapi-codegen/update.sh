#!/bin/bash

fix_external_ref_bug() {
    target=$1
    import_line=$2
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

for conf_name in *.conf; do
    schema_name="../schemas/${conf_name%%.*}.yaml"
    echo "$schema_name..."
    oapi-codegen -config "$conf_name" "$schema_name"
done

# note: the code below exists because of this issue:
# https://github.com/deepmap/oapi-codegen/issues/377

targets=(
    "../go/client/authorizer/authorizer.go"
    "../go/client/client_service/client_service.go"
    "../go/client/dashboard/dashboard.go"
    "../go/client/discovery/discovery.go"
    "../go/client/federation/federation.go"
    "../go/client/license_service/license_service.go"
    "../go/client/project_service/project_service.go"
    "../go/client/user_service/user_service.go"
    "../go/client/tunnel_mgmt/tunnel.go"
)

for target in "${targets[@]}"; do
    fix_external_ref_bug "$target" "import ("
done