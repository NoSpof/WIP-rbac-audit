#!/bin/bash
while [[ true ]]
do 
    date
    echo "Get Roles"
    /kubectl  get roles --all-namespaces -o json > /data/roles.json
    echo "Get Cluster Roles"
    /kubectl  get clusterroles -o json > /data/clusterroles.json
    echo "Get Roles Bindings"
    /kubectl  get rolebindings --all-namespaces -o json > /data/rolebindings.json
    echo "Get Cluster Roles Bindings"
    /kubectl  get clusterrolebindings -o json > /data/clusterrolebindings.json
    echo "sleep 30"
    echo "##########################"
    sleep 30

done