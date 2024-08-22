# STIR/SHAKEN Certificate Repository Compliance

## Clearwave Communications

Name: `https://cdn-cr.cgah.tnsi.com/certs/8c047baba4b6c6f319541b73856679bda03e801f`\
Tested At: 22 Aug 24 15:18 UTC\
Time: 22ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 22 Aug 24 16:06 UTC