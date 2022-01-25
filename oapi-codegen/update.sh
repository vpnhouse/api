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

### fix external ref bug for the federation-client
target=../go/client/federation/federation.go
ln=$(grep -m1 -n "github.com/deepmap/oapi-codegen/pkg/runtime" $target | cut -d: -f1)
ed $target << END
${ln}a
	externalRef1 "github.com/Codename-Uranium/api/go/server/common"
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

	externalRef1 "github.com/Codename-Uranium/api/go/server/common"
.
w
q
END

sed -i -e 's/*Error/*externalRef1.Error/g' $target
sed -i -e 's/var dest Error/var dest externalRef1.Error/g' $target
