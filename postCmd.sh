#!/bin/bash
apt update && apt install python3 python3.13-venv -y
python3 -m venv .venv
source .venv/bin/activate
pip install pre-commit
