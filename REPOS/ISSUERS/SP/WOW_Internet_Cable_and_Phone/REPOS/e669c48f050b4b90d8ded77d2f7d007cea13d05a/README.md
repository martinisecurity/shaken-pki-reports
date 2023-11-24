# STIR/SHAKEN Certificate Repository Compliance

## WOW Internet Cable and Phone

Name: `https://cdn-cr.cgah.tnsi.com/certs/086231fe3b29598f23fa70a5dd994ed7153040a1`\
Tested At: 24 Nov 23 11:00 UTC\
Time: 12ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 24 Nov 23 11:17 UTC