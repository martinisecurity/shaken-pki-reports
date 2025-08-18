# STIR/SHAKEN Certificate Repository Compliance

## Aeneas Communications

Name: `https://cdn-cr.cgah.tnsi.com/certs/16b0b2ffa0adc4c8507782e9c76c8d1c6e4b71e5`\
Tested At: 18 Aug 25 20:05 UTC\
Time: 20ms

### Issues

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_cache_header](../../ISSUES/e_atis_cache_header/README.md) | error | ATIS-1000074 | The STI-VS shall implement the cache behavior. The Cache-Control header is missed |
| [w_atis_content_type](../../ISSUES/w_atis_content_type/README.md) | warn | ATIS-1000080 | HTTP response should contain Content-Type header and it's value should be application/pem-certificate-chain |

Generated: 18 Aug 25 21:13 UTC