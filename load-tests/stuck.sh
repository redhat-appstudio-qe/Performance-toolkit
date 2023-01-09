
#!/bin/bash
set -e
set -o pipefail
set -u

# osascript -e 'tell app "Terminal"
#     do script "oc proxy"
# end tell' 

namespaces=$(oc get ns --selector='toolchain.dev.openshift.com/type=appstudio' -o name)


for i in $namespaces
do
    SUBSTRING=$(echo $i| cut -d'/' -f 2)
    echo "deleting namespace : $SUBSTRING"
    oc get namespace $SUBSTRING -o json > tmp.json
    jq -c '.spec.finalizers = []' tmp.json > tmp.$$.json && mv tmp.$$.json tmp.json
    curl -k -H "Content-Type: application/json" -X PUT --data-binary @tmp.json http://127.0.0.1:8001/api/v1/namespaces/$SUBSTRING/finalize
done
