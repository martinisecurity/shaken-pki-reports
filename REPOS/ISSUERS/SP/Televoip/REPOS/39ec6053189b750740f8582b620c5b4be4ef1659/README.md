# STIR/SHAKEN Certificate Repository Compliance

## Televoip

Name: `https://ssc.getsipnav.com/certs/525bf6e144190a8232ed2c15770df226099ff1ff`\
Tested At: 16 Mar 23 19:09 UTC\
Time: 74ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 16 Mar 23 19:18 UTC