import test from 'node:test';import assert from 'node:assert/strict';import {readFile} from 'node:fs/promises';import {filterCollection,gatewayOrigin,gatewayForTarget,substitute,boundedResponse,validateWorkspace} from './workbench.mjs';
 test('Public source has no embedded message or collection exporter',async()=>{const source=await readFile(new URL('./workbench.mjs',import.meta.url),'utf8');assert.ok(!source.includes('postmanCollection'));assert.ok(!source.includes('SYNTHETIC-001'));assert.ok(!source.includes('createObjectURL'));const html=await readFile(new URL('./workbench.html',import.meta.url),'utf8');assert.ok(!html.includes('id="download"'));assert.ok(html.includes('id="workspace-file"'));});
 test('Foreign or empty workspaces are rejected',()=>{assert.throws(()=>validateWorkspace({},'expected'));});
 test('Missing variables block execution; revision is numeric',()=>{assert.throws(()=>substitute('/{{id}}',{}));assert.deepEqual(substitute({revision:'{{revision}}'},{revision:'2'}),{revision:2});});
 test('Unsafe credential destinations rejected',()=>{assert.throws(()=>gatewayOrigin('http://example.com'));assert.throws(()=>gatewayOrigin('https://user:pass@example.com'));assert.equal(gatewayOrigin('http://127.0.0.1:8791').origin,'http://127.0.0.1:8791');});
 test('Responses bounded',async()=>{await assert.rejects(()=>boundedResponse(new Response('x'.repeat(1048577))));});

test("Address Intelligence never falls back to the payment service",()=>{assert.equal(gatewayForTarget({target:"address"},"http://127.0.0.1:8793","http://127.0.0.1:8794").port,"8794");assert.equal(gatewayForTarget({},"http://127.0.0.1:8793","http://127.0.0.1:8794").port,"8793");assert.throws(()=>gatewayForTarget({target:"address"},"http://127.0.0.1:8793",""));});

test('Two selectors isolate collections and flows',()=>{const items=[{collection:'SEPA',flow:'SDD Core'},{collection:'SEPA',flow:'SCT'},{collection:'CBPR+',flow:'MX'}];assert.deepEqual(filterCollection(items,'SEPA','SDD Core'),[items[0]]);assert.deepEqual(filterCollection(items,'CBPR+','MX'),[items[2]]);});
test('SEPA requires its explicitly selected rail gateway',()=>{assert.equal(gatewayForTarget({target:'sepa',rail:'TIPS'},'http://127.0.0.1:8793','',{TIPS:'http://127.0.0.1:8797'}).port,'8797');assert.throws(()=>gatewayForTarget({target:'sepa',rail:'RT1'},'http://127.0.0.1:8793','',{}));});
test('RTGS and CHAPS use their own gateways without payment-service fallback',()=>{
 const rails={T2:'http://127.0.0.1:8798',BOE_RTGS:'http://127.0.0.1:8799',CHAPS:'http://127.0.0.1:8800'};
 for(const [rail,url] of Object.entries(rails))assert.equal(gatewayForTarget({target:'rail',rail},'http://127.0.0.1:8793','',rails).origin,url);
 assert.throws(()=>gatewayForTarget({target:'rail',rail:'CHAPS'},'http://127.0.0.1:8793','',{}));
});

test('Workspace import rejects a graph belonging to another project',()=>{const workspace={key:'expected',api:[{name:'Read',method:'GET',path:'/record',role:'accessToken'}],mcpRequests:[],messages:{mt103:'',pacs008:'',pacs002:''},vars:{},graph:{version:1,project:'foreign',nodes:[],edges:[],evidence:[],limits:[],scope:'Snapshot'}};assert.throws(()=>validateWorkspace(workspace,'expected'));});
