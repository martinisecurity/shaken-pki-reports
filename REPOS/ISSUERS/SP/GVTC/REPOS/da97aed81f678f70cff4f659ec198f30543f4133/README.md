# STIR/SHAKEN Certificate Repository Compliance

## GVTC

Name: `https://cdn-cr.cgah.tnsi.com/certs/90b8c01a8cc0c06e82b9517d599d24c14dbb859b`\
Tested At: 27 Nov 23 23:10 UTC\
Time: 13ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 27 Nov 23 23:28 UTC