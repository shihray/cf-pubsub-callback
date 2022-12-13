#!/bin/bash
set -x

source .cloudbuild/script/loadenv.sh

gcloud --project="${PROJECT_ID}" functions deploy --region=asia-east1 ${RUNMODE}-pan-avatar-transformation \
  --service-account ${CF_SERVICE_ACCOUNT} \
  --entry-point ${ENTRYPOINT} --runtime go116 \
  --trigger-http \
  --env-vars-file .cloudbuild/config/${BRANCH_NAME}.env.yaml \
  --vpc-connector "serverless-pan" \
  --memory 512
#  --trigger-resource "${TRIGGER_RESOURCE}" \
#  --trigger-event "google.storage.object.finalize" \

#gcloud --project="${STORAGE_PROJECT_ID}" functions deploy --region=asia-east1 ${RUNMODE}-pan-avatar-delegate \
#  --service-account ${CF_SERVICE_ACCOUNT} \
#  --entry-point ${STORAGE_ENTRYPOINT} --runtime go116 \
#  --trigger-event "google.storage.object.finalize" \
#  --trigger-resource "${TRIGGER_RESOURCE}" \
#  --memory 128
