# STIR/SHAKEN Certificate Repository Compliance

## Segra

Name: `https://cdn-cr.cgah.tnsi.com/certs/86728e9104b39e125eb3c8a4f8a8224325991d00`\
Tested At: 04 Oct 24 15:31 UTC\
Time: 76ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 04 Oct 24 16:29 UTC