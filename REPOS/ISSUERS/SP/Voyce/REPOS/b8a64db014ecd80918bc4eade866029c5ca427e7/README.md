# STIR/SHAKEN Certificate Repository Compliance

## Voyce

Name: `https://ssc.getsipnav.com/certs/b00badc474bbf0d965554e422647a4fcc426eb0c`\
Tested At: 11 Jan 23 23:10 UTC\
Time: 84ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 11 Jan 23 23:18 UTC