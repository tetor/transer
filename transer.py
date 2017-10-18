#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import sys, json, yaml, fire


def transer(in_file):
    print(json2yaml(in_file))

def json2yaml(in_file):
    with open(in_file) as f:
        return yaml.safe_dump(json.load(f), default_flow_style=False)

if __name__ == '__main__':
    fire.Fire(transer)
