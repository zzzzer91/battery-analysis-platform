#!/bin/bash

python3 -m celery -A task worker --concurrency=2