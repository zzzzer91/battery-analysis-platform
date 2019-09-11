from typing import Dict

import yaml


def load(path: str) -> Dict:
    with open(path) as f:
        return yaml.load(f, Loader=yaml.FullLoader)


app_conf = load('/conf/app.yml')['celery']
