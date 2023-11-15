# STIR/SHAKEN Certificate Repository Compliance

## Consolidated Telcom ND

Name: `https://cdn-cr.cgah.tnsi.com/certs/1e9efaedff05b3fc118ca5c6c413a10ce0a84b98`\
Tested At: 15 Nov 23 17:58 UTC\
Time: 9ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 15 Nov 23 18:10 UTC