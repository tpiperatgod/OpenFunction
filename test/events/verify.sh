#!/bin/bash

#
# Copyright 2022 The OpenFunction Authors.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
# http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#

export KUBECONFIG=/tmp/e2e-k8s.config

MANIFESTS_PATH="./test/events/manifests.yaml"

if [ "${DOCKERHUB_REPO_PREFIX}" ]; then
  declare -a FUNCTIONS=("e2e-v1beta1-cron-in-kafka-out")
  for fn in "${FUNCTIONS[@]}";do
    IMAGE_NAME=${DOCKERHUB_REPO_PREFIX}/$(NAME=${fn} yq e 'select(.metadata.name == env(NAME) ).spec.image' ${MANIFESTS_PATH}|awk -F\/ '{ print $2 }') NAME=${fn} yq -i e 'select(.metadata.name == env(NAME)).spec.image |= env(IMAGE_NAME)' ${MANIFESTS_PATH}
  done
fi

kubectl apply -f ${MANIFESTS_PATH} > /dev/null 2>&1

kubectl logs $(kubectl get po -l openfunction.io/serving=$(kubectl get functions sink-a -o jsonpath='{.status.serving.resourceRef}') \
  -o jsonpath='{.items[0].metadata.name}') function |grep "Hello" > /dev/null 2>&1
if [ $? -eq 0 ]; then
  kubectl logs $(kubectl get po -l openfunction.io/serving=$(kubectl get functions sink-b -o jsonpath='{.status.serving.resourceRef}') \
    -o jsonpath='{.items[0].metadata.name}') function |grep "Hello" > /dev/null 2>&1
  if [ $? -eq 0 ]; then
    echo "events: success" | yq
    kubectl delete -f ${MANIFESTS_PATH} > /dev/null 2>&1
  fi
fi
