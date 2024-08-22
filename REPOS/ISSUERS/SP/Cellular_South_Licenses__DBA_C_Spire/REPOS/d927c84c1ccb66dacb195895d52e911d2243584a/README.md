# STIR/SHAKEN Certificate Repository Compliance

## Cellular South Licenses  DBA C Spire

Name: `https://cdn-cr.cgah.tnsi.com/certs/d1af372b9cdb3ceba1f85f4ec3657797b3eb1c7f`\
Tested At: 22 Aug 24 15:16 UTC\
Time: 10ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 15:44 UTC