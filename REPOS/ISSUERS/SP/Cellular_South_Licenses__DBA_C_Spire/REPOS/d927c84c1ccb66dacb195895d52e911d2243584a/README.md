# STIR/SHAKEN Certificate Repository Compliance

## Cellular South Licenses  DBA C Spire

Name: `https://cdn-cr.cgah.tnsi.com/certs/d1af372b9cdb3ceba1f85f4ec3657797b3eb1c7f`\
Tested At: 04 Oct 24 15:31 UTC\
Time: 62ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC