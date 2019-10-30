#!/bin/bash

ABSOLUTE_PROJECT_PATH="$(cd "$(dirname "$0")/.." || exit;pwd)"

FRONTEND_DIR="${ABSOLUTE_PROJECT_PATH}/frontend"
RESOURCE_DIR="${ABSOLUTE_PROJECT_PATH}/resource"

cd "${FRONTEND_DIR}" || exit
npm run build || exit

cd "$ABSOLUTE_PROJECT_PATH" || exit

if [ ! -d "${RESOURCE_DIR}" ]; then
  mkdir "${RESOURCE_DIR}"
fi

if [ -d "${RESOURCE_DIR}/dist" ]; then
  rm -r "${RESOURCE_DIR}/dist"
fi

mv "${FRONTEND_DIR}/dist" "${RESOURCE_DIR}/dist"|| exit

echo '前端编译完成！'