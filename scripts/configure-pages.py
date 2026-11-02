#!/usr/bin/env python3
"""Apply the public core endpoint to a Pages deployment; no credentials are accepted."""
import pathlib,json,os,urllib.parse
p=pathlib.Path('docs/client-config.json');config=json.loads(p.read_text());value=os.environ.get('CORE_API_URL','').strip()
if value:
 u=urllib.parse.urlsplit(value)
 if u.scheme!='https' or not u.hostname or u.username or u.password or u.query or u.fragment:raise SystemExit('CORE_API_URL must be an HTTPS API address without credentials.')
 config['apiBaseUrl']=value.rstrip('/')
p.write_text(json.dumps(config,indent=2)+'\n')
