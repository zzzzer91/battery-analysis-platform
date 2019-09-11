#!/bin/bash

python3 -m celery -A task worker --loglevel=INFO --concurrency=2