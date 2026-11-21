#!/usr/bin/env python3
"""End-to-end classroom scenarios; requires Docker Compose and Python 3."""
import json
import subprocess
import time
import urllib.error
import urllib.parse
import urllib.request
from pathlib import Path
ROOT = Path(__file__).resolve().parents[1]

def request(port, path, params=None, body=None):
    url = f'http://127.0.0.1:{port}{path}'
    if params:
        url += '?' + urllib.parse.urlencode(params)
    req = urllib.request.Request(url, data=body)
    if body is not None:
        req.add_header('Content-Type', 'application/xml')
    try:
        with urllib.request.urlopen(req, timeout=5) as response:
            return response.status, response.read().decode()
    except urllib.error.HTTPError as error:
        return error.code, error.read().decode()

def check(name, port):
    def get(path, **params):
        return request(port, path, params)
    if name == 'sqli':
        payload = "' OR '1'='1' -- "
        assert len(json.loads(get('/vulnerable/users', name='alice')[1])) == 1
        assert len(json.loads(get('/safe/users', name='alice')[1])) == 1
        assert len(json.loads(get('/vulnerable/users', name=payload)[1])) == 3
        assert json.loads(get('/safe/users', name=payload)[1]) == []
    elif name == 'command-injection':
        payload = "Alice'; id; #"
        assert get('/vulnerable/greet', name='Alice') == (200, 'Alice\n')
        assert 'uid=' in get('/vulnerable/greet', name=payload)[1]
        assert get('/safe/greet', name=payload) == (200, payload+'\n')
    elif name == 'xxe':
        normal = (ROOT/name/'examples/normal.xml').read_bytes()
        attack = (ROOT/name/'examples/xxe.xml').read_bytes()
        for mode in ['vulnerable', 'safe']:
            assert request(port, f'/{mode}/xml', body=normal) == (200, 'Hello XML')
        assert 'DEMO_SECRET=' in request(port, '/vulnerable/xml', body=attack)[1]
        assert request(port, '/safe/xml', body=attack)[0] == 400
    elif name == 'ssrf':
        assert 'Public catalog' in get('/vulnerable/import', url='http://127.0.0.1:9090/public')[1]
        assert 'DEMO_INTERNAL_TOKEN=' in get('/vulnerable/import', url='http://127.0.0.1:9090/admin')[1]
        assert 'Public catalog' in get('/safe/import', document='catalog')[1]
        assert get('/safe/import', url='http://127.0.0.1:9090/admin')[0] == 400
    else:
        for mode in ['vulnerable', 'safe']:
            assert 'Public lecture' in get(f'/{mode}/files', name='hello.txt')[1]
        assert 'DEMO_SECRET=' in get('/vulnerable/files', name='../secret.txt')[1]
        assert get('/safe/files', name='../secret.txt')[0] == 403

for name, port in [('sqli',8081),('command-injection',8082),('xxe',8083),('ssrf',8084),('path-traversal',8085)]:
    directory = ROOT/name
    def compose(*args, **kwargs):
        return subprocess.run(['docker','compose',*args], cwd=directory, **kwargs)
    try:
        compose('up','--build','-d', check=True)
        for attempt in range(60):
            try:
                request(port, '/')  # 404 also proves the HTTP listener is ready.
                break
            except (OSError, urllib.error.URLError):
                time.sleep(1)
        else:
            raise RuntimeError(f'{name}: server not ready')
        check(name, port)
        print(f'PASS {name}', flush=True)
    except BaseException:
        compose('logs','--no-color')
        raise
    finally:
        compose('down','-v', check=True)
