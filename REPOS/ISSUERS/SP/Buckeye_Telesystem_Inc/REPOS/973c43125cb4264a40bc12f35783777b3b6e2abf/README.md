# STIR/SHAKEN Certificate Repository Compliance

## Buckeye Telesystem Inc

Name: `https://cdn-cr.cgah.tnsi.com/certs/3c6ee1503f2b739a7b68a95572cba9a30bfad75f`\
Tested At: 22 Aug 24 15:16 UTC\
Time: 31ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 15:44 UTC