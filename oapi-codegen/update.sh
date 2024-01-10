#!/bin/sh

oc=oapi-codegen

ls *.conf |
while read conf_name
do
  schema_name="../schemas/${conf_name%%.*}.yaml"
  echo "$schema_name..."
  $oc -config $conf_name $schema_name
done

# note: the code below exists because of this issue:
# https://github.com/deepmap/oapi-codegen/issues/377

### fix external ref bug for the authorizer-client
target=../go/client/authorizer/authorizer.go
ln=$(grep -m1 -n "github.com/deepmap/oapi-codegen/pkg/runtime" $target | cut -d: -f1)
ed $target << END
${ln}a
	externalRef1 "github.com/vpnhouse/api/go/server/common"
.
w
q
END

sed -i -e 's/*Error/*externalRef1.Error/g' $target
sed -i -e 's/var dest Error/var dest externalRef1.Error/g' $target

### fix external ref bug for the federation-client
target=../go/client/federation/federation.go
ln=$(grep -m1 -n "github.com/deepmap/oapi-codegen/pkg/runtime" $target | cut -d: -f1)
ed $target << END
${ln}a
	externalRef1 "github.com/vpnhouse/api/go/server/common"
.
w
q
END

sed -i -e 's/*Error/*externalRef1.Error/g' $target
sed -i -e 's/var dest Error/var dest externalRef1.Error/g' $target

### fix external ref bug for the user_service-client
target=../go/client/user_service/user_service.go
ln=$(grep -m1 -n "github.com/deepmap/oapi-codegen/pkg/runtime" $target | cut -d: -f1)
ed $target << END
${ln}a
	externalRef1 "github.com/vpnhouse/api/go/server/common"
.
w
q
END

sed -i -e 's/*Error/*externalRef1.Error/g' $target
sed -i -e 's/var dest Error/var dest externalRef1.Error/g' $target

### fix external ref bug for the client_service-client
target=../go/client/client_service/client_service.go
ln=$(grep -m1 -n "import (" $target | cut -d: -f1)
ed $target << END
${ln}a
	externalRef1 "github.com/vpnhouse/api/go/server/common"
.
w
q
END

sed -i -e 's/*Error/*externalRef1.Error/g' $target
sed -i -e 's/var dest Error/var dest externalRef1.Error/g' $target

### fix external ref bug for the project_service-client
target=../go/client/project_service/project_service.go
ln=$(grep -m1 -n "github.com/deepmap/oapi-codegen/pkg/runtime" $target | cut -d: -f1)
ed $target << END
${ln}a
	externalRef1 "github.com/vpnhouse/api/go/server/common"
.
w
q
END

sed -i -e 's/*Error/*externalRef1.Error/g' $target
sed -i -e 's/var dest Error/var dest externalRef1.Error/g' $target

### fix external ref bug for the license_service-client
target=../go/client/license_service/license_service.go
ln=$(grep -m1 -n "github.com/deepmap/oapi-codegen/pkg/runtime" $target | cut -d: -f1)
ed $target << END
${ln}a
	externalRef1 "github.com/vpnhouse/api/go/server/common"
.
w
q
END

sed -i -e 's/*Error/*externalRef1.Error/g' $target
sed -i -e 's/var dest Error/var dest externalRef1.Error/g' $target

### fix external ref bug in the tunnel_mgmt-client
target=../go/client/tunnel_mgmt/tunnel.go
ln=$(grep -m1 -n "strings" $target | cut -d: -f1)
ed ../go/client/tunnel_mgmt/tunnel.go << END
${ln}a

	externalRef1 "github.com/vpnhouse/api/go/server/common"
.
w
q
END

sed -i -e 's/*Error/*externalRef1.Error/g' $target
sed -i -e 's/var dest Error/var dest externalRef1.Error/g' $target
