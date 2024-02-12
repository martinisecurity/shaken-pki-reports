# STIR/SHAKEN Certificate Repository Compliance

## Avid Communication

Name: `https://cdn-cr.cgah.tnsi.com/certs/1b0b321bad320b960a4f0fc8f2408d0daa110730`\
Tested At: 12 Feb 24 16:26 UTC\
Time: 17ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 12 Feb 24 17:02 UTC