# STIR/SHAKEN Certificate Repository Compliance

## GCI

Name: `https://cdn-cr.cgah.tnsi.com/certs/69fae8e574f06fb796e08c7b6ea847f2563cea5b`\
Tested At: 21 Nov 23 17:37 UTC\
Time: 22ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 21 Nov 23 17:53 UTC