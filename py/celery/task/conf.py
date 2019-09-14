import os
from typing import Dict

import yaml


def load(path: str) -> Dict:
    with open(path) as f:
        return yaml.load(f, Loader=yaml.SafeLoader)


app_conf = load(os.getenv('CONF_FILE'))['celery']
