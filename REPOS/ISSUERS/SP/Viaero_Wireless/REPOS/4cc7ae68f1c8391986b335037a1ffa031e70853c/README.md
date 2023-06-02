# STIR/SHAKEN Certificate Repository Compliance

## Viaero Wireless

Name: `https://cdn-cr.cgah.tnsi.com/certs/836dbe7b4ab9ad475219697d1553bed65e3d3463`\
Tested At: 02 Jun 23 01:00 UTC\
Time: 9ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 02 Jun 23 01:12 UTC